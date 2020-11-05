package model

type Task struct {
	Type string `json:"type"`
	Spec struct {
		DataSchema struct {
			DataSource    string `json:"dataSource"`
			TimestampSpec struct {
				Column       string      `json:"column"`
				Format       string      `json:"format"`
				MissingValue interface{} `json:"missingValue"`
			} `json:"timestampSpec"`
			DimensionsSpec struct {
				Dimensions []struct {
					Type               string `json:"type"`
					Name               string `json:"name"`
					MultiValueHandling string `json:"multiValueHandling"`
					CreateBitmapIndex  bool   `json:"createBitmapIndex"`
				} `json:"dimensions"`
				DimensionExclusions []string `json:"dimensionExclusions"`
			} `json:"dimensionsSpec"`
			MetricsSpec []struct {
				Type       string      `json:"type"`
				Name       string      `json:"name"`
				FieldName  string      `json:"fieldName,omitempty"`
				Expression interface{} `json:"expression,omitempty"`
			} `json:"metricsSpec"`
			GranularitySpec struct {
				Type               string      `json:"type"`
				SegmentGranularity string      `json:"segmentGranularity"`
				QueryGranularity   string      `json:"queryGranularity"`
				Rollup             bool        `json:"rollup"`
				Intervals          interface{} `json:"intervals"`
			} `json:"granularitySpec"`
			TransformSpec struct {
				Filter     interface{}   `json:"filter"`
				Transforms []interface{} `json:"transforms"`
			} `json:"transformSpec"`
		} `json:"dataSchema"`
		IoConfig struct {
			Type        string `json:"type"`
			InputSource struct {
				Type    string        `json:"type"`
				BaseDir string        `json:"baseDir"`
				Filter  string        `json:"filter"`
				Files   []interface{} `json:"files"`
			} `json:"inputSource"`
			InputFormat struct {
				Type                  string        `json:"type"`
				Columns               []interface{} `json:"columns"`
				ListDelimiter         interface{}   `json:"listDelimiter"`
				FindColumnsFromHeader bool          `json:"findColumnsFromHeader"`
				SkipHeaderRows        int           `json:"skipHeaderRows"`
			} `json:"inputFormat"`
			AppendToExisting bool `json:"appendToExisting"`
		} `json:"ioConfig"`
		TuningConfig struct {
			Type              string      `json:"type"`
			MaxRowsPerSegment int         `json:"maxRowsPerSegment"`
			MaxRowsInMemory   int         `json:"maxRowsInMemory"`
			MaxBytesInMemory  int         `json:"maxBytesInMemory"`
			MaxTotalRows      interface{} `json:"maxTotalRows"`
			NumShards         interface{} `json:"numShards"`
			SplitHintSpec     interface{} `json:"splitHintSpec"`
			PartitionsSpec    struct {
				Type              string      `json:"type"`
				MaxRowsPerSegment int         `json:"maxRowsPerSegment"`
				MaxTotalRows      interface{} `json:"maxTotalRows"`
			} `json:"partitionsSpec"`
			IndexSpec struct {
				Bitmap struct {
					Type                       string `json:"type"`
					CompressRunOnSerialization bool   `json:"compressRunOnSerialization"`
				} `json:"bitmap"`
				DimensionCompression string      `json:"dimensionCompression"`
				MetricCompression    string      `json:"metricCompression"`
				LongEncoding         string      `json:"longEncoding"`
				SegmentLoader        interface{} `json:"segmentLoader"`
			} `json:"indexSpec"`
			IndexSpecForIntermediatePersists struct {
				Bitmap struct {
					Type                       string `json:"type"`
					CompressRunOnSerialization bool   `json:"compressRunOnSerialization"`
				} `json:"bitmap"`
				DimensionCompression string      `json:"dimensionCompression"`
				MetricCompression    string      `json:"metricCompression"`
				LongEncoding         string      `json:"longEncoding"`
				SegmentLoader        interface{} `json:"segmentLoader"`
			} `json:"indexSpecForIntermediatePersists"`
			MaxPendingPersists           int           `json:"maxPendingPersists"`
			ForceGuaranteedRollup        bool          `json:"forceGuaranteedRollup"`
			ReportParseExceptions        bool          `json:"reportParseExceptions"`
			PushTimeout                  int           `json:"pushTimeout"`
			SegmentWriteOutMediumFactory interface{}   `json:"segmentWriteOutMediumFactory"`
			MaxNumConcurrentSubTasks     int           `json:"maxNumConcurrentSubTasks"`
			MaxRetry                     int           `json:"maxRetry"`
			TaskStatusCheckPeriodMs      int           `json:"taskStatusCheckPeriodMs"`
			ChatHandlerTimeout           string        `json:"chatHandlerTimeout"`
			ChatHandlerNumRetries        int           `json:"chatHandlerNumRetries"`
			MaxNumSegmentsToMerge        int           `json:"maxNumSegmentsToMerge"`
			TotalNumMergeTasks           int           `json:"totalNumMergeTasks"`
			LogParseExceptions           bool          `json:"logParseExceptions"`
			MaxParseExceptions           int64         `json:"maxParseExceptions"`
			MaxSavedParseExceptions      int           `json:"maxSavedParseExceptions"`
			BuildV9Directly              bool          `json:"buildV9Directly"`
			PartitionDimensions          []interface{} `json:"partitionDimensions"`
		} `json:"tuningConfig"`
	} `json:"spec"`
}
