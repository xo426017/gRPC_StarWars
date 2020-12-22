using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Grpc.Core;
using Main;

namespace cstarwars
{
  public class StarwarsService : starwars.starwarsBase
  {
    private readonly Dictionary<int, Character> _characters;

    private readonly Dictionary<Episode, List<Review>> _data = new Dictionary<Episode, List<Review>>();

    public StarwarsService()
    {
      _characters = CreateCharacters().ToDictionary(t => t.Id);
    }

    public override Task<Character> GetHero(GetHeroRequest req, ServerCallContext context)
    {
      Console.WriteLine("Incoming get hero request: Episode = {0} " + req.Episode);

      var response = req.Episode == Episode.Empire ? _characters[1000] : _characters[2001];

      return Task.FromResult(response);
    }

    public override Task<SearchCharacterResponse> SearchCharacter(SearchCharacterRequest req, ServerCallContext context)
    {
      Console.WriteLine("Incoming search character request: name = {0} ", req.Name);

      var resp = new SearchCharacterResponse();
      foreach (var character in _characters.Values
        .Where(t => (!string.IsNullOrEmpty(req.Name) && t.Name.Contains(req.Name, StringComparison.OrdinalIgnoreCase)) || t.Id == req.Id))
      {
        resp.Characters.Add(character);
      }

      return Task.FromResult(resp);
    }

    public override Task<Review> AddReview(Review review, ServerCallContext context)
    {
      Console.WriteLine("Incoming add review request: {0} : {1}, {2} ", review.Episode, review.Stars, review.Commentary);

      if (!_data.TryGetValue(review.Episode, out List<Review> reviews))
      {
        reviews = new List<Review>();
        _data[review.Episode] = reviews;
      }

      reviews.Add(review);

      return Task.FromResult(review);
    }

    public override Task<GetReviewsResponse> GetReviews(GetReviewsRequest req, ServerCallContext context)
    {
      Console.WriteLine("Incoming get reviews request: Episode = {0} ", req.Episode);

      var resp = new GetReviewsResponse();
      if (_data.TryGetValue(req.Episode, out List<Review> reviews))
      {
        resp.Reviews.AddRange(reviews);
      }

      return Task.FromResult(resp);
    }

    private static IEnumerable<Character> CreateCharacters()
    {
      yield return new Character
      {
        Id = 1000,
        Name = "Luke Skywalker",
        Friends = { "Han Solo", "Leia Organa", "C-3PO", "R2-D2" },
        AppearsIn = { Episode.NewHope, Episode.Empire, Episode.Jedi },
        HomePlanet = "Tatooine"
      };

      yield return new Character
      {
        Id = 1001,
        Name = "Darth Vader",
        Friends = { "Wilhuff Tarkin" },
        AppearsIn = { Episode.NewHope, Episode.Empire, Episode.Jedi },
        HomePlanet = "Tatooine"
      };

      yield return new Character
      {
        Id = 1002,
        Name = "Han Solo",
        Friends = { "Luke Skywalker", "Leia Organa", "C-3PO" },
        AppearsIn = { Episode.NewHope, Episode.Empire, Episode.Jedi },
      };

      yield return new Character
      {
        Id = 1003,
        Name = "Leia Organa",
        Friends = { "Han Solo", "Luke Skywalker", "C-3PO", "R2-D2" },
        AppearsIn = { Episode.NewHope, Episode.Empire, Episode.Jedi },
        HomePlanet = "Alderaan"
      };

      yield return new Character
      {
        Id = 1004,
        Name = "Wilhuff Tarkin",
        Friends = { "Darth Vader" },
        AppearsIn = { Episode.NewHope }
      };

      yield return new Character
      {
        Id = 2000,
        Name = "C-3PO",
        Friends = { "Han Solo", "Luke Skywalker", "Leia Organa", "R2-D2" },
        AppearsIn = { Episode.NewHope, Episode.Empire, Episode.Jedi },
        PrimaryFunction = "Protocol"
      };

      yield return new Character
      {
        Id = 2001,
        Name = "R2-D2",
        Friends = { "Han Solo", "Luke Skywalker", "C-3PO" },
        AppearsIn = { Episode.NewHope, Episode.Empire, Episode.Jedi },
        PrimaryFunction = "Astromech"
      };
    }
  }
}
