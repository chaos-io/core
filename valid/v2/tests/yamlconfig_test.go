package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chaos-io/core/httputil/headers"
	"github.com/chaos-io/core/ptr"
	inspection2 "github.com/chaos-io/core/valid/v2/inspection"
	rule2 "github.com/chaos-io/core/valid/v2/rule"
)

const (
	sizeSeparator rune = 'x'
)

type yamlConfig struct {
	Resizers []yamlResizer
	Storages []yamlStorage
	Buckets  []yamlBucket
}

func (y yamlConfig) Validate() error {
	return v2.valid.Struct(&y,
		v2.valid.Value(&y.Resizers, rule2.NotEmpty, rule2.Unique(func(v *inspection2.Inspected) string {
			return v.Interface.(yamlResizer).Name
		})),
		v2.valid.Value(&y.Storages, rule2.NotEmpty, rule2.Unique(func(v *inspection2.Inspected) string {
			return v.Interface.(yamlStorage).Name
		})),
		v2.valid.Value(&y.Buckets, rule2.NotEmpty, rule2.Unique(func(v *inspection2.Inspected) string {
			buk := v.Interface.(yamlBucket)
			if buk.Target != "" {
				return buk.Target
			}
			return buk.Path
		})),
	)
}

type yamlResizer struct {
	Name           string
	BaseURL        string
	RequestTimeout time.Duration
	RetryCount     int
	Priority       int
}

func (y yamlResizer) Validate() error {
	return v2.valid.Struct(&y,
		v2.valid.Value(&y.Name, rule2.NotEmpty, rule2.IsASCII),
		v2.valid.Value(&y.BaseURL, rule2.NotEmpty, rule2.IsURL),
		v2.valid.Value(&y.RequestTimeout, rule2.OmitEmpty(rule2.IsPositive)),
		v2.valid.Value(&y.RetryCount, rule2.OmitEmpty(rule2.IsPositive)),
		v2.valid.Value(&y.Priority, rule2.OmitEmpty(rule2.IsPositive)),
	)
}

type yamlStorage struct {
	Name    string
	BaseURL string
}

func (y yamlStorage) Validate() error {
	return v2.valid.Struct(&y,
		v2.valid.Value(&y.Name, rule2.NotEmpty, rule2.IsASCII),
		v2.valid.Value(&y.BaseURL, rule2.NotEmpty, rule2.IsURL),
	)
}

type yamlBucket struct {
	Path     string
	Target   string
	Storage  string
	Sizes    []string
	Aliases  map[string]string
	Features yamlFeatures
}

func (y yamlBucket) Validate() error {
	mAliasMessage := "must contain alias with key 'm'"

	return v2.valid.Struct(&y,
		v2.valid.Value(&y.Path, rule2.NotEmpty, rule2.IsASCII, rule2.IsAbsDir),
		v2.valid.Value(&y.Target, rule2.OmitEmpty(rule2.IsASCII, rule2.IsAbsDir)),
		v2.valid.Value(&y.Sizes, rule2.NotEmpty, rule2.Unique(rule2.ValueAsKey), rule2.Each(rule2.Is2DMeasurements("x"))),
		v2.valid.Value(&y.Aliases,
			rule2.Message(mAliasMessage, rule2.NotEmpty, rule2.HasKey("m")),
			rule2.Each(rule2.Is2DMeasurements("x")),
		),
		v2.valid.Value(&y.Features),
	)
}

type yamlFeatures struct {
	AllowProcessing     *bool
	BackgroundColor     *string
	ConvertTo           *string
	PreferType          *string
	ColorPalette        *uint8
	Watermark           *int
	DisableWatermarkFor []string
	Quality             *uint8
	ProxyExtensions     []string
	FallbackImage       *string
	EnableWMDK          *bool
}

func (y yamlFeatures) Validate() error {
	validImageType := rule2.InSlice([]string{
		string(headers.TypeImageJPEG),
		string(headers.TypeImageGIF),
		string(headers.TypeImagePNG),
		string(headers.TypeImageWebP),
		string(headers.TypeImageSVG),
	})

	return v2.valid.Struct(&y,
		v2.valid.Value(&y.BackgroundColor, rule2.OmitEmpty(rule2.IsHexColor)),
		v2.valid.Value(&y.ConvertTo, rule2.OmitEmpty(validImageType)),
		v2.valid.Value(&y.PreferType, rule2.OmitEmpty(validImageType)),
		v2.valid.Value(&y.ColorPalette, rule2.OmitEmpty(rule2.InRange(0, 128))),
		v2.valid.Value(&y.Watermark, rule2.OmitEmpty(rule2.IsPositive)),
		v2.valid.Value(&y.DisableWatermarkFor, rule2.OmitEmpty(rule2.Each(rule2.NotEmpty))),
		v2.valid.Value(&y.Quality, rule2.OmitEmpty(rule2.InRange(0, 100))),
		v2.valid.Value(&y.ProxyExtensions, rule2.OmitEmpty(validImageType)),
		v2.valid.Value(&y.FallbackImage, rule2.OmitEmpty(rule2.IsURL)),
	)
}

func TestValidate_YamlConfig(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		c := yamlConfig{
			Resizers: []yamlResizer{
				{
					Name:           "shakalat",
					BaseURL:        "http://shakalat.local/",
					RequestTimeout: time.Second,
					RetryCount:     1,
					Priority:       100,
				},
				{
					Name:           "mds",
					BaseURL:        "http://resize.mds.local/",
					RequestTimeout: 2 * time.Second,
					RetryCount:     1,
					Priority:       10,
				},
			},
			Storages: []yamlStorage{
				{
					Name:    "s3",
					BaseURL: "http://s3.mds.yandex.ru/my-bucket/",
				},
			},
			Buckets: []yamlBucket{
				{
					Path:    "/test/valid_path/",
					Target:  "",
					Storage: "s3",
					Sizes: []string{
						"200x200",
						"400x400",
					},
					Aliases: map[string]string{
						"m": "200x200",
						"l": "400x400",
					},
					Features: yamlFeatures{
						BackgroundColor: ptr.String("FFFFFF"),
						PreferType:      ptr.String(string(headers.TypeImageWebP)),
						Quality:         ptr.Uint8(90),
					},
				},
			},
		}

		err := c.Validate()
		assert.NoError(t, err)
	})

	t.Run("empty_top_level", func(t *testing.T) {
		c := yamlConfig{
			Resizers: []yamlResizer{
				{
					Name:           "shakalat",
					BaseURL:        "http://shakalat.local/",
					RequestTimeout: time.Second,
					RetryCount:     1,
					Priority:       100,
				},
				{
					Name:           "мдс",
					BaseURL:        "http://resize.mds.local/",
					RequestTimeout: 2 * time.Second,
					RetryCount:     -1,
					Priority:       10,
				},
			},
			Storages: nil,
			Buckets: []yamlBucket{
				{
					Path:    "/test/valid_path",
					Target:  "",
					Storage: "s3",
					Sizes: []string{
						"200x200",
						"400x400",
					},
					Aliases: map[string]string{
						"l": "400x400",
					},
					Features: yamlFeatures{
						BackgroundColor: ptr.String("JKLMN"),
						PreferType:      ptr.String(string(headers.TypeImageWebP)),
						Quality:         ptr.Uint8(90),
					},
				},
			},
		}

		ic := inspection2.Inspect(c)                      // inspected config
		ir := inspection2.Inspect(c.Resizers[1])          // inspected resizer
		ib := inspection2.Inspect(c.Buckets[0])           // inspected bucket
		ift := inspection2.Inspect(c.Buckets[0].Features) // inspected feature

		expected := rule2.Errors{
			rule2.NewFieldError(
				&ic.Fields[0].Field,
				rule2.NewFieldError(&ir.Fields[0].Field, rule2.ErrInvalidCharacters),
			),
			rule2.NewFieldError(
				&ic.Fields[0].Field,
				rule2.NewFieldError(&ir.Fields[3].Field, rule2.ErrNegativeValue),
			),
			rule2.NewFieldError(&ic.Fields[1].Field, rule2.ErrEmptyValue),
			rule2.NewFieldError(
				&ic.Fields[2].Field,
				rule2.NewFieldError(&ib.Fields[0].Field, rule2.ErrPatternMismatch),
			),
			rule2.NewFieldError(
				&ic.Fields[2].Field,
				rule2.NewFieldError(&ib.Fields[4].Field, &rule2.MessageErr{
					Msg: "must contain alias with key 'm'",
					Err: rule2.ErrUnexpected,
				}),
			),
			rule2.NewFieldError(
				&ic.Fields[2].Field,
				rule2.NewFieldError(
					&ib.Fields[5].Field,
					rule2.NewFieldError(&ift.Fields[1].Field, rule2.ErrInvalidStringLength),
				),
			),
		}
		assert.Equal(t, expected, c.Validate())
	})
}

func BenchmarkValidate_YamlConfig(b *testing.B) {
	c := yamlConfig{
		Resizers: []yamlResizer{
			{
				Name:           "shakalat",
				BaseURL:        "http://shakalat.local/",
				RequestTimeout: time.Second,
				RetryCount:     1,
				Priority:       100,
			},
			{
				Name:           "mds",
				BaseURL:        "http://resize.mds.local/",
				RequestTimeout: 2 * time.Second,
				RetryCount:     1,
				Priority:       10,
			},
		},
		Storages: []yamlStorage{
			{
				Name:    "s3",
				BaseURL: "http://s3.mds.yandex.ru/my-bucket/",
			},
		},
		Buckets: []yamlBucket{
			{
				Path:    "/retailers/icons/",
				Target:  "",
				Storage: "s3",
				Sizes: []string{
					"100x100",
					"200x200",
					"400x400",
				},
				Aliases: map[string]string{
					"s": "100x100",
					"m": "200x200",
					"l": "400x400",
				},
				Features: yamlFeatures{
					PreferType: ptr.String(string(headers.TypeImageWebP)),
				},
			},
			{
				Path:    "/items/",
				Target:  "/offers/",
				Storage: "s3",
				Sizes: []string{
					"350x350",
					"450x450",
					"900x900",
				},
				Aliases: map[string]string{
					"s": "350x350",
					"m": "450x450",
					"l": "900x900",
				},
				Features: yamlFeatures{
					BackgroundColor:     ptr.String("FFFFFF"),
					ConvertTo:           ptr.String(string(headers.TypeImageJPEG)),
					Quality:             ptr.Uint8(90),
					DisableWatermarkFor: []string{".nwm"},
				},
			},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.Validate()
	}
}
