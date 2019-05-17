// Copyright © 2018 One Concern

package cmd

import (
	"log"
	"time"

	"github.com/oneconcern/datamon/pkg/dlogger"

	"github.com/oneconcern/datamon/pkg/core"
	"github.com/oneconcern/datamon/pkg/storage/gcs"
	"github.com/oneconcern/datamon/pkg/storage/localfs"
	"github.com/spf13/afero"

	"github.com/spf13/cobra"
)

// Mount a read only view of a bundle
var mountBundleCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount a bundle",
	Long:  "Mount a readonly, non-interactive view of the entire data that is part of a bundle",
	Run: func(cmd *cobra.Command, args []string) {

		DieIfNotAccessible(bundleOptions.DataPath)

		metadataSource, err := gcs.New(repoParams.MetadataBucket, config.Credential)
		if err != nil {
			logFatalln(err)
		}
		blobStore, err := gcs.New(repoParams.BlobBucket, config.Credential)
		if err != nil {
			logFatalln(err)
		}
		consumableStore := localfs.New(afero.NewBasePathFs(afero.NewOsFs(), bundleOptions.DataPath))

		err = setLatestBundle(metadataSource)
		if err != nil {
			logFatalln(err)
		}
		bd := core.NewBDescriptor()
		bundle := core.New(bd,
			core.Repo(repoParams.RepoName),
			core.BundleID(bundleOptions.ID),
			core.BlobStore(blobStore),
			core.ConsumableStore(consumableStore),
			core.MetaStore(metadataSource),
		)
		logger, err := dlogger.GetLogger(logLevel)
		if err != nil {
			log.Fatalln("Failed to set log level:" + err.Error())
		}
		fs, err := core.NewReadOnlyFS(bundle, logger)
		if err != nil {
			logFatalln(err)
		}
		err = fs.MountReadOnly(bundleOptions.MountPath)
		if err != nil {
			logFatalln(err)
		}
		for {
			time.Sleep(time.Hour)
		}
	},
}

func init() {

	requiredFlags := []string{addRepoNameOptionFlag(mountBundleCmd)}
	addBucketNameFlag(mountBundleCmd)
	addBlobBucket(mountBundleCmd)
	addBundleFlag(mountBundleCmd)
	addLogLevel(mountBundleCmd)
	// todo: #165 add --cpuprof to all commands via root
	addCPUProfFlag(mountBundleCmd)
	requiredFlags = append(requiredFlags, addDataPathFlag(mountBundleCmd))
	requiredFlags = append(requiredFlags, addMountPathFlag(mountBundleCmd))

	for _, flag := range requiredFlags {
		err := mountBundleCmd.MarkFlagRequired(flag)
		if err != nil {
			logFatalln(err)
		}
	}

	bundleCmd.AddCommand(mountBundleCmd)
}
