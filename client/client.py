#!/usr/bin/python3

from pathlib import Path
import sys

sys.path.append(str(Path(__file__).parent / '../'))

from dir_watcher import dir_watcher_pb2, dir_watcher_pb2_grpc
import grpc

import argparse

class Client:
    def __init__(self, port):
        self.port = port
        self.channel = grpc.insecure_channel(f"127.0.0.1:{self.port}")
        self.grpc_client = dir_watcher_pb2_grpc.DirWatcherStub(self.channel)

    def getDirContents(self, dir):
        try:
            result = self.grpc_client.Do(dir_watcher_pb2.DirWatchRequest(directory=dir))
            for r in result.content:
                if r != "":
                    print(r)
        except grpc.RpcError as grpc_error:
            print(grpc_error)
        
if __name__ == "__main__":
    parser = argparse.ArgumentParser(prog = 'Client')
    parser.add_argument('-p', '--port', default=5030)
    parser.add_argument('-d', '--dir', default=".")
    args = parser.parse_args()
    client = Client(args.port)
    client.getDirContents(args.dir)
