package ipfs

import (
	"fmt"
	"io"
	"log"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

// IPFSClient represents a simple client to interact with an IPFS node
type IPFSClient struct {
	sh *shell.Shell
}

// NewIPFSClient creates a new instance of IPFSClient
func NewIPFSClient(nodeAddress string) *IPFSClient {
	return &IPFSClient{
		sh: shell.NewShell(nodeAddress),
	}
}

// AddFile adds a file to IPFS and returns the CID.
func (c *IPFSClient) AddFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	cid, err := c.sh.Add(file)
	if err != nil {
		return "", fmt.Errorf("failed to add file to IPFS: %v", err)
	}
	return cid, nil
}

// GetFile retrieves a file from IPFS using its CID and saves it to the specified output path.
func (c *IPFSClient) GetFile(cid, outputPath string) error {
	reader, err := c.sh.Cat(cid)
	if err != nil {
		return fmt.Errorf("failed to retrieve file from IPFS: %v", err)
	}
	defer reader.Close()

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, reader)
	if err != nil {
		return fmt.Errorf("failed to write file content: %v", err)
	}

	log.Printf("File saved to %s\n", outputPath)
	return nil
}