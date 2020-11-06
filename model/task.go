package model

type Task struct {
	Type string `json:"type"`
	Spec Spec   `json:"spec"`
}

type TimestampSpec struct {
	Column       string      `json:"column"`
	Format       string      `json:"format"`
	MissingValue interface{} `json:"missingValue"`
}

type Dimensions struct {
	Type               string  `json:"type"`
	Name               string  `json:"name"`
	MultiValueHandling *string `json:"multiValueHandling"`
	CreateBitmapIndex  *bool   `json:"createBitmapIndex"`
}

type DimensionsSpec struct {
	Dimensions          []Dimensions `json:"dimensions"`
	DimensionExclusions []string     `json:"dimensionExclusions"`
}

type MetricsSpec struct {
	Type       string      `json:"type"`
	Name       string      `json:"name"`
	FieldName  string      `json:"fieldName,omitempty"`
	Expression interface{} `json:"expression,omitempty"`
}

type GranularitySpec struct {
	Type               string   `json:"type"`
	SegmentGranularity string   `json:"segmentGranularity"`
	QueryGranularity   string   `json:"queryGranularity"`
	Rollup             bool     `json:"rollup"`
	Intervals          []string `json:"intervals"`
}

type TransformSpec struct {
	Filter     interface{}   `json:"filter"`
	Transforms []interface{} `json:"transforms"`
}

type DataSchema struct {
	DataSource      string          `json:"dataSource"`
	TimestampSpec   TimestampSpec   `json:"timestampSpec"`
	DimensionsSpec  DimensionsSpec  `json:"dimensionsSpec"`
	MetricsSpec     []MetricsSpec   `json:"metricsSpec"`
	GranularitySpec GranularitySpec `json:"granularitySpec"`
	TransformSpec   TransformSpec   `json:"transformSpec"`
}

type InputSource struct {
	Type    string        `json:"type"`
	BaseDir string        `json:"baseDir"`
	Filter  string        `json:"filter"`
	Files   []interface{} `json:"files"`
}

type InputFormat struct {
	Type                  string        `json:"type"`
	Columns               []interface{} `json:"columns"`
	ListDelimiter         interface{}   `json:"listDelimiter"`
	FindColumnsFromHeader bool          `json:"findColumnsFromHeader"`
	SkipHeaderRows        int           `json:"skipHeaderRows"`
}

type IoConfig struct {
	Type             string      `json:"type"`
	InputSource      InputSource `json:"inputSource"`
	InputFormat      InputFormat `json:"inputFormat"`
	AppendToExisting bool        `json:"appendToExisting"`
}

type PartitionsSpec struct {
	Type                 string  `json:"type"`
	TargetRowsPerSegment int     `json:"targetRowsPerSegment"`
	PartitionDimension   *string `json:"partitionDimension"`
	MaxRowsPerSegment    *int    `json:"maxRowsPerSegment"`
}

type Bitmap struct {
	Type                       string `json:"type"`
	CompressRunOnSerialization bool   `json:"compressRunOnSerialization"`
}

type IndexSpec struct {
	Bitmap               Bitmap      `json:"bitmap"`
	DimensionCompression string      `json:"dimensionCompression"`
	MetricCompression    string      `json:"metricCompression"`
	LongEncoding         string      `json:"longEncoding"`
	SegmentLoader        interface{} `json:"segmentLoader"`
}

type IndexSpecForIntermediatePersists struct {
	Bitmap               Bitmap      `json:"bitmap"`
	DimensionCompression string      `json:"dimensionCompression"`
	MetricCompression    string      `json:"metricCompression"`
	LongEncoding         string      `json:"longEncoding"`
	SegmentLoader        interface{} `json:"segmentLoader"`
}

type TuningConfig struct {
	Type                  string         `json:"type"`
	PartitionsSpec        PartitionsSpec `json:"partitionsSpec"`
	ForceGuaranteedRollup *bool          `json:"forceGuaranteedRollup"`
}

type Spec struct {
	DataSchema   DataSchema    `json:"dataSchema"`
	IoConfig     IoConfig      `json:"ioConfig"`
	TuningConfig *TuningConfig `json:"tuningConfig"`
}
