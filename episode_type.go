package podcast

// PodcastType specifies the episode type.
//
// Its values can be one of the following:
//
// if an episode is a trailer or bonus content, use this tag.
// Where the episodeType value can be one of the following:
//
// Full (default). Specify full when you are submitting the complete content of your show.
// Trailer. Specify trailer when you are submitting a short, promotional piece of content that represents a preview of your current show.
// Bonus. Specify bonus when you are submitting extra content for your show (for example, behind the scenes information or interviews with the cast) or cross-promotional content for another show.
//
// The rules for using trailer and bonus tags depend on whether the <itunes:season> and <itunes:episode> tags have values:
// Trailer:
// No season or episode number: a show trailer
// A season number and no episode number: a season trailer. (Note: an episode trailer should have a different <guid> than the actual episode)
// Episode number and optionally a season number: an episode trailer/teaser, later replaced with the actual episode
//
// Bonus:
// No season or episode number: a show bonus
// A season number: a season bonus
// Episode number and optionally a season number: a bonus episode related to a specific episode

const (
	Full EpisodeType = iota
	Trailer
	Bonus
)

const (
	episodeTypeDefault = "Full"
)

// EpisodeType specifies the episode type.
//
// Its values can be one of the following:
//
// Episodic (default). Specify episodic when episodes are intended to be
// consumed without any specific order. Apple Podcasts will present newest
// episodes first and display the publish date (required) of each episode.
// If organized into seasons, the newest season will be presented first
// - otherwise, episodes will be grouped by year published, newest first.
//
// For new subscribers, Apple Podcasts adds the newest, most recent episode
// in their Library.
//
// Serial. Specify serial when episodes are intended to be consumed in
// sequential order. Apple Podcasts will present the oldest episodes
// first and display the episode numbers (required) of each episode. If
// organized into seasons, the newest season will be presented first and
// <itunes:episode> numbers must be given for each episode.
//
// For new subscribers, Apple Podcasts adds the first episode to their
// Library, or the entire current season if using seasons.
type EpisodeType int

// String returns the MIME type encoding of the specified EnclosureType.
func (et EpisodeType) String() string {
	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	switch et {
	case Full:
		return "Full"
	case Trailer:
		return "Trailer"
	case Bonus:
		return "Bonus"
	}
	return episodeTypeDefault
}
