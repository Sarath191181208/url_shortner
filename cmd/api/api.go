package api

import (
	"sarath/url_shortner/internal/cache"
	"sarath/url_shortner/internal/data"
	"sarath/url_shortner/internal/json/logger"
)


type Application struct{
  Logger logger.ApplicationLogger
  Db *data.Models
  Cache cache.Cache
}

