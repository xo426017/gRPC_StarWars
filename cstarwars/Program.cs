using System;
using Grpc.Core;
using Main;

namespace cstarwars
{
  public class Program
  {
    public static void Main(string[] args)
    {
      // Build the server
      Console.WriteLine("Starting C# server on port 9002");
      Server server = new Server {
          Services = { starwars.BindService(new StarwarsService()) },
          Ports = { new ServerPort("localhost", 9002, ServerCredentials.Insecure)}
      };
      server.Start();

      // Call the Go server on port 9001
      // Set up gRPC client
      Channel channel = new Channel("localhost:9001", ChannelCredentials.Insecure);
      var client = new starwars.starwarsClient(channel);

      Console.WriteLine("Enter the command to call the Go server...");
      Console.WriteLine("   c/name- search characters");
      Console.WriteLine("   h/name - get hero by episode");
      Console.WriteLine("   k/name - create review");
      Console.WriteLine("   r/name - get reviews by episode");
      Console.WriteLine("   x - exit");

      for (; ; )
      {
        var command = Console.ReadLine();
        var s = command.Split("/");

        if (s[0].Equals("x")) break;

        Episode episode;
        switch (s[0])
        {
          case "c":
            var resp = client.SearchCharacter(new SearchCharacterRequest {Name = s[1]});
            Console.WriteLine("characters: {0}", resp);
            break;
          case "h":
            Enum.TryParse(s[1], out episode);
            var hero = client.GetHero(new GetHeroRequest {Episode = episode});
            Console.WriteLine("hero: {0}", hero);
            break;
          case "k":
            if (s.Length != 4) break;

            Enum.TryParse(s[1], out episode);
            int.TryParse(s[2], out int stars);
            var review = new Review {Stars = stars, Commentary = s[3], Episode = episode};
            var revResp = client.AddReview(review);
            Console.WriteLine("review: {0}", revResp);
            break;
          case "r":
            Enum.TryParse(s[1], out episode);
            var reviewResp = client.GetReviews(new GetReviewsRequest {Episode = episode });
            Console.WriteLine("reviews: {0}", reviewResp);
            break;
        }
      }

      // Block for server termination
      Console.ReadKey();
      channel.ShutdownAsync().Wait();
      server.ShutdownAsync().Wait();
    }
  }
}
