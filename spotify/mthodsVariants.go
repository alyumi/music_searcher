package spotify

import (
	"github.com/alyumi/music_searcher/spotify/methods"
)

type Options struct {
	Track  bool
	Artist bool
	Album  bool
}

type Methods struct {
	options *Options
}

type Setter func(*Options)

func WithTrack(is bool) Setter {
	return func(opts *Options) {
		opts.Track = is
	}
}

func WithArtist(is bool) Setter {
	return func(opts *Options) {
		opts.Artist = is
	}
}

func WithAlbum(is bool) Setter {
	return func(opts *Options) {
		opts.Album = is
	}
}

func AddStruct(setters ...Setter) *Methods {
	options := &Options{
		Track:  false,
		Album:  false,
		Artist: false,
	}

	for _, set := range setters {
		set(options)
	}

	return &Methods{
		options: options,
	}
}

func ImplementStructs(method Methods) {
	opts := method.options
	_ = opts
	// switch opts {
	// case opts.Track == true:
	// 	implementTrack()
	// }
}

func implementTrack() *methods.Track {
	return &methods.Track{}
}

func implementArtist() *methods.Artist {
	return &methods.Artist{}
}

func implementAlbum() *methods.Album {
	return &methods.Album{}
}
