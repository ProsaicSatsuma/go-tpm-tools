package cmd

import (
	"fmt"

	"github.com/ProsaicSatsuma/go-tpm-tools/client"
	"github.com/google/go-tpm/tpm2"
	"github.com/google/go-tpm/tpmutil"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read <pcr>",
	Short: "Read from the TPM",
	Long:  `Read from the TPM`,
	Args:  cobra.NoArgs,
}

var pcrCmd = &cobra.Command{
	Use:   "pcr",
	Short: "Read PCRs from the TPM",
	Long: `Read PCRs from the TPM

Based on the --pcrs and --hash-algo flags, this reads the contents of the TPM's
PCRs for that hash algorithm.

If --pcrs is not provided, all pcrs are read for that hash algorithm.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		rwc, err := openTpm()
		if err != nil {
			return err
		}
		defer rwc.Close()

		sel := getSelection()
		if len(sel.PCRs) == 0 {
			sel = client.FullPcrSel(hashAlgo)
		}

		fmt.Fprintf(debugOutput(), "Reading pcrs (%v)\n", sel.PCRs)
		pcrList, err := client.ReadPCRs(rwc, sel)
		if err != nil {
			return err
		}

		for idx := uint32(0); idx < client.NumPCRs; idx++ {
			if val, ok := pcrList.Pcrs[idx]; ok {
				if _, err = fmt.Fprintf(dataOutput(), "%2d: %x\n", idx, val); err != nil {
					return err
				}
			}
		}
		return nil
	},
}

var nvReadCmd = &cobra.Command{
	Use:   "nvdata",
	Short: "Read TPM NVData",
	Long: `Read NVData at a particular NVIndex

Based on the --index flag, this reads all of the NVData present at that NVIndex.
The read is authenticated with the owner hierarchy and an empty password.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		rwc, err := openTpm()
		if err != nil {
			return err
		}
		defer rwc.Close()

		data, err := tpm2.NVReadEx(rwc, tpmutil.Handle(nvIndex), tpm2.HandleOwner, "", 0)
		if err != nil {
			return err
		}
		if _, err := dataOutput().Write(data); err != nil {
			return fmt.Errorf("cannot output NVData: %w", err)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(readCmd)
	readCmd.AddCommand(pcrCmd)
	readCmd.AddCommand(nvReadCmd)
	addOutputFlag(pcrCmd)
	addPCRsFlag(pcrCmd)
	addHashAlgoFlag(pcrCmd)
	addIndexFlag(nvReadCmd)
	nvReadCmd.MarkPersistentFlagRequired("index")
	addOutputFlag(nvReadCmd)
}
