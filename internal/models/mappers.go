package models

import (
	gamepb "github.com/viktoralyoshin/playhub-proto/gen/go/games" // Твой сгенерированный пакет

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToProto конвертирует твою доменную модель в формат gRPC ответа
func (g *Game) ToProto() *gamepb.Game {
	return &gamepb.Game{
		Id:               g.ID,
		IgdbId:           g.IGDBID,
		Name:             g.Name,
		Slug:             g.Slug,
		Summary:          g.Summary,
		Rating:           g.Rating,
		Hypes:            uint32(g.Hypes), // Кастим int в uint32
		FirstReleaseDate: g.FirstReleaseDate,
		ReleaseDates:     g.ReleaseDates,
		CoverUrl:         g.CoverURL,
		ArtworkUrls:      g.ArtworkURLs,
		Screenshots:      g.Screenshots,
		Genres:           g.Genres,
		Themes:           g.Themes,
		Platforms:        g.Platforms,

		// Конвертация времени Go -> Protobuf Timestamp
		CreatedAt: timestamppb.New(g.CreatedAt),
		UpdatedAt: timestamppb.New(g.UpdatedAt),
	}
}

// FromProto создает доменную модель из gRPC сообщения
func GameFromProto(pb *gamepb.Game) *Game {
	return &Game{
		ID:               pb.Id,
		IGDBID:           pb.IgdbId,
		Name:             pb.Name,
		Slug:             pb.Slug,
		Summary:          pb.Summary,
		Rating:           pb.Rating,
		Hypes:            int(pb.Hypes),
		FirstReleaseDate: pb.FirstReleaseDate,
		ReleaseDates:     pb.ReleaseDates,
		CoverURL:         pb.CoverUrl,
		ArtworkURLs:      pb.ArtworkUrls,
		Screenshots:      pb.Screenshots,
		Genres:           pb.Genres,
		Themes:           pb.Themes,
		Platforms:        pb.Platforms,

		// Конвертация Protobuf Timestamp -> Go Time
		CreatedAt: pb.CreatedAt.AsTime(),
		UpdatedAt: pb.UpdatedAt.AsTime(),
	}
}
