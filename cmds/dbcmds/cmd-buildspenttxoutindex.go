package dbcmds

import (
	"github.com/WikiLeaksFreedomForce/local-blockchain-parser/blockdb"
)

type BuildSpentTxOutIndexCommand struct {
	dbFile     string
	datFileDir string
	startBlock uint64
	endBlock   uint64
}

func NewBuildSpentTxOutIndexCommand(startBlock, endBlock uint64, datFileDir, dbFile string) *BuildSpentTxOutIndexCommand {
	return &BuildSpentTxOutIndexCommand{
		dbFile:     dbFile,
		datFileDir: datFileDir,
		startBlock: startBlock,
		endBlock:   endBlock,
	}
}

func (cmd *BuildSpentTxOutIndexCommand) RunCommand() error {
	db, err := blockdb.NewBlockDB(cmd.dbFile, cmd.datFileDir)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.IndexDATFileSpentTxOuts(cmd.startBlock, cmd.endBlock)
	if err != nil {
		return err
	}
	return nil
}
