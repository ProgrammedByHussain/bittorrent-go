# BitTorrent Client in Go

A lightweight BitTorrent client implementation written in Go that can download files from torrent networks.

## Features

- Parse `.torrent` files using bencode encoding
- Connect to tracker servers to discover peers
- Establish peer connections with BitTorrent handshake protocol
- Download pieces in parallel from multiple peers
- Verify piece integrity using SHA-1 hashes
- Handle BitTorrent protocol messages (choke, unchoke, interested, have, bitfield, request, piece)
- Concurrent download workers for efficient file retrieval

## Installation

```bash
# Clone the repository
git clone https://github.com/ProgrammedByHussain/bittorrent-go.git
cd bittorrent-go

# Install dependencies
go mod download
```

## Usage

```bash
go run . <torrent-file> <output-file>
```

**Example:**
```bash
go run . ubuntu.torrent ubuntu.iso
```

## Project Structure

```
.
├── bitfield/          # Bitfield operations for tracking piece availability
├── client/            # TCP client for peer connections
├── handshake/         # BitTorrent handshake protocol implementation
├── message/           # Protocol message encoding/decoding
├── p2p/               # Peer-to-peer download orchestration
├── peers/             # Peer data structures and parsing
├── torrentfile/       # Torrent file parsing and tracker communication
└── main.go            # Entry point
```

## How It Works

1. **Parse Torrent File**: Reads and decodes the `.torrent` file to extract metadata including tracker URL, piece hashes, and file information
2. **Request Peers**: Contacts the tracker server to get a list of peers sharing the file
3. **Establish Connections**: Performs handshakes with available peers
4. **Download Pieces**: Spawns concurrent workers to download file pieces from multiple peers
5. **Verify Integrity**: Validates each downloaded piece against its SHA-1 hash
6. **Assemble File**: Combines verified pieces into the complete file

## Technical Details

- **Piece Size**: Standard 262 KB pieces with configurable block size (16 KB)
- **Concurrency**: Multiple goroutines handle parallel downloads from different peers
- **Protocol**: Implements core BitTorrent protocol v1.0 messages
- **Timeout Handling**: 3-second handshake timeout, 30-second piece download timeout

## Dependencies

- `github.com/jackpal/bencode-go` - Bencode encoding/decoding

## Requirements

- Go 1.25.1 or higher

## License

This project is for educational purposes.
