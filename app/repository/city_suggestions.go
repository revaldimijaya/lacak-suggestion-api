package repository

import (
	"bufio"
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func (r *Repository) LoadDataCities(ctx context.Context, req LoadDataRequest) (err error) {
	file, err := os.Open(req.DataPath)
	if err != nil {
		return errors.Wrap(err, "error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "\t")
		if len(fields) < 6 {
			continue
		}

		name := fields[1]
		lat, err1 := strconv.ParseFloat(fields[4], 64)
		lon, err2 := strconv.ParseFloat(fields[5], 64)
		if err1 != nil || err2 != nil {
			continue
		}

		r.DataSource.Insert(City{Name: name, Latitude: lat, Longitude: lon})
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "error scanning file")
	}

	return
}
