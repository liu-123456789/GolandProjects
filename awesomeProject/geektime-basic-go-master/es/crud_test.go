package es

import (
	"context"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
	"time"
)

type ElasticSearchTestSuite struct {
	suite.Suite
	es *elastic.Client
}

func (s *ElasticSearchTestSuite) SetupSuite() {
	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	require.NoError(s.T(), err)
	s.es = client
}

func (s *ElasticSearchTestSuite) TestCreateIndex() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	// 这是一个链式调用，你可以通过链式调用来构造复杂请求。
	// 重复创建会报错，所以你可以换一个名字
	resp, err := s.es.Indices.Create("user_idx_test",
		s.es.Indices.Create.WithContext(ctx),
		s.es.Indices.Create.WithBody(strings.NewReader(`
{  
  "settings": {  
    "number_of_shards": 3,  
    "number_of_replicas": 2  
  },  
  "mappings": {  
    "properties": {
      "email": {  
        "type": "text"  
      },  
      "phone": {  
        "type": "keyword"  
      },  
      "birthday": {  
        "type": "date"  
      }
    }  
  }  
}
`)))
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 200, resp.StatusCode)
}

func (s *ElasticSearchTestSuite) TestPutDoc() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := s.es.Index("user_idx_test", strings.NewReader(`
{  
  "email": "john@example.com",  
  "phone": "1234567890",  
  "birthday": "2000-01-01"  
}
`), s.es.Index.WithContext(ctx))
	require.NoError(s.T(), err)
	require.NotNil(s.T(), resp)
	assert.Equal(s.T(), 201, resp.StatusCode)
}

func (s *ElasticSearchTestSuite) TestGetDoc() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := s.es.Search(s.es.Search.WithContext(ctx),
		s.es.Search.WithBody(strings.NewReader(`
{  
  "query": {  
    "range": {  
      "birthday": {
        "gte": "1990-01-01"
      }
    }  
  }  
}
`)))
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 200, resp.StatusCode)
}

func TestElasticSearch(t *testing.T) {
	suite.Run(t, new(ElasticSearchTestSuite))
}
