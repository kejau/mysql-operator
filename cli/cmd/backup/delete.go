package backup

import (
	"fmt"
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"

	"github.com/grtl/mysql-operator/cli/cmd/config"
	"github.com/grtl/mysql-operator/cli/cmd/interact"
	"github.com/grtl/mysql-operator/cli/cmd/util"
	"github.com/grtl/mysql-operator/operator/backup"
)

var removePVC bool

var backupDeleteCmd = &cobra.Command{
	Use:   "delete [backup name]",
	Short: "A short description of backup delete",
	Long: `A longer description of backup delete with examples:
msp backup delete "my-cluster"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		options := util.ExtractOptions(cmd)

		for _, arg := range args {
			err := deleteBackup(arg, options)
			util.FailOnErrorOrForceContinue(err, options)
		}
	},
}

func deleteBackup(backupName string, options *util.Options) error {
	fmt.Printf("You are trying to remove MySQL Backup: %s/%s\n", options.Namespace, backupName)
	answer, _ := interact.YesNoInput(options)
	if !answer {
		return nil
	}

	backupsInterface := config.GetConfig().Clientset().CrV1().MySQLBackups(options.Namespace)
	err := backupsInterface.Delete(backupName, &v1.DeleteOptions{})

	if removePVC && (err == nil || options.Force) {
		deleteErr := deletePVC(clusterName, options)
		return errors.NewAggregate([]error{err, deleteErr})
	}

	return err
}

func deletePVC(backupName string, options *util.Options) error {
	fmt.Printf("You are trying to remove PVC for MySQL Backup: %s/%s\n", options.Namespace, clusterName)
	answer, _ := interact.YesNoInput(options)
	if !answer {
		return nil
	}

	pvcInterface := config.GetConfig().KubeClientset().CoreV1().PersistentVolumeClaims(options.Namespace)
	return pvcInterface.Delete(backup.PVCName(backupName), &v1.DeleteOptions{})
}

func init() {
	backupDeleteCmd.PersistentFlags().BoolVarP(&removePVC, "remove-pvc", "r", false, "remove PersistentVolumeClaim along with the backup")
	Cmd.AddCommand(backupDeleteCmd)
}
