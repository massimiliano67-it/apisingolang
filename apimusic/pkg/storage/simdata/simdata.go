package simdata

import (
	"context"
	"fmt"
	"sync"

	music "apimusic/pkg/music"
)

type musicRepository struct {
	mtx    sync.RWMutex
	musics map[string]*music.Music
}

func NewMusicRepository(musics map[string]*music.Music) music.MusicRepository {
	if musics == nil {
		musics = make(map[string]*music.Music)
	}

	return &musicRepository{
		musics: musics,
	}
}

func (r *musicRepository) CreateMusic(ctx context.Context, g *music.Music) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if err := r.checkIfExists(g.ID); err != nil {
		return err
	}
	r.musics[g.ID] = g
	return nil
}

func (r *musicRepository) FetchMusics(ctx context.Context) ([]*music.Music, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]*music.Music, 0, len(r.musics))
	for _, value := range r.musics {
		values = append(values, value)
	}
	return values, nil
}

func (r *musicRepository) DeleteMusic(ctx context.Context, ID string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.musics, ID)

	return nil
}

func (r *musicRepository) UpdateMusic(ctx context.Context, ID string, g *music.Music) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.musics[ID] = g
	return nil
}

func (r *musicRepository) FetchMusicByID(ctx context.Context, ID string) (*music.Music, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.musics {
		if v.ID == ID {
			return v, nil
		}
	}

	return nil, fmt.Errorf("The ID %s doesn't exist", ID)
}

func (r *musicRepository) checkIfExists(ID string) error {
	for _, v := range r.musics {
		if v.ID == ID {
			return fmt.Errorf("The music %s is already exist", ID)
		}
	}

	return nil
}
