import grpc
import livescore.score_pb2_grpc as pb2_grpc
import livescore.score_pb2 as pb2


class LiveScoreClient(object):
    """
    Client for Live Score
    """

    def __init__(self):
        self.host = 'localhost'
        self.server_port = 50004

        # instantiate a channel
        self.channel = grpc.insecure_channel(
            '{}:{}'.format(self.host, self.server_port))

        # bind the client and the server
        self.stub = pb2_grpc.ScoreServiceStub(self.channel)

    def get_live_score(self, country):
        """
        Client function to call the rpc for GetServerResponse
        """
        matchRequest = pb2.ListMatchesRequest(country=country)
        print(f'{matchRequest}')
        return self.stub.ListMatches(matchRequest)


if __name__ == '__main__':
    client = LiveScoreClient()
    result = client.get_live_score(country="India")
    print(f'{result}')