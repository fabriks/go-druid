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
	Type               string `json:"type"`
	Name               string `json:"name"`
	MultiValueHandling string `json:"multiValueHandling"`
	CreateBitmapIndex  bool   `json:"createBitmapIndex"`
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
	Type               string      `json:"type"`
	SegmentGranularity string      `json:"segmentGranularity"`
	QueryGranularity   string      `json:"queryGranularity"`
	Rollup             bool        `json:"rollup"`
	Intervals          interface{} `json:"intervals"`
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
	Type              string      `json:"type"`
	MaxRowsPerSegment int         `json:"maxRowsPerSegment"`
	MaxTotalRows      interface{} `json:"maxTotalRows"`
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
	Type                             string                           `json:"type"`
	MaxRowsPerSegment                int                              `json:"maxRowsPerSegment"`
	MaxRowsInMemory                  int                              `json:"maxRowsInMemory"`
	MaxBytesInMemory                 int                              `json:"maxBytesInMemory"`
	MaxTotalRows                     interface{}                      `json:"maxTotalRows"`
	NumShards                        interface{}                      `json:"numShards"`
	SplitHintSpec                    interface{}                      `json:"splitHintSpec"`
	PartitionsSpec                   PartitionsSpec                   `json:"partitionsSpec"`
	IndexSpec                        IndexSpec                        `json:"indexSpec"`
	IndexSpecForIntermediatePersists IndexSpecForIntermediatePersists `json:"indexSpecForIntermediatePersists"`
	MaxPendingPersists               int                              `json:"maxPendingPersists"`
	ForceGuaranteedRollup            bool                             `json:"forceGuaranteedRollup"`
	ReportParseExceptions            bool                             `json:"reportParseExceptions"`
	PushTimeout                      int                              `json:"pushTimeout"`
	SegmentWriteOutMediumFactory     interface{}                      `json:"segmentWriteOutMediumFactory"`
	MaxNumConcurrentSubTasks         int                              `json:"maxNumConcurrentSubTasks"`
	MaxRetry                         int                              `json:"maxRetry"`
	TaskStatusCheckPeriodMs          int                              `json:"taskStatusCheckPeriodMs"`
	ChatHandlerTimeout               string                           `json:"chatHandlerTimeout"`
	ChatHandlerNumRetries            int                              `json:"chatHandlerNumRetries"`
	MaxNumSegmentsToMerge            int                              `json:"maxNumSegmentsToMerge"`
	TotalNumMergeTasks               int                              `json:"totalNumMergeTasks"`
	LogParseExceptions               bool                             `json:"logParseExceptions"`
	MaxParseExceptions               int64                            `json:"maxParseExceptions"`
	MaxSavedParseExceptions          int                              `json:"maxSavedParseExceptions"`
	BuildV9Directly                  bool                             `json:"buildV9Directly"`
	PartitionDimensions              []interface{}                    `json:"partitionDimensions"`
}
type Spec struct {
	DataSchema   DataSchema   `json:"dataSchema"`
	IoConfig     IoConfig     `json:"ioConfig"`
	TuningConfig TuningConfig `json:"tuningConfig"`
}
