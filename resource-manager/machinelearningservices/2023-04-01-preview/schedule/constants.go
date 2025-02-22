package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoDeleteCondition string

const (
	AutoDeleteConditionCreatedGreaterThan      AutoDeleteCondition = "CreatedGreaterThan"
	AutoDeleteConditionLastAccessedGreaterThan AutoDeleteCondition = "LastAccessedGreaterThan"
)

func PossibleValuesForAutoDeleteCondition() []string {
	return []string{
		string(AutoDeleteConditionCreatedGreaterThan),
		string(AutoDeleteConditionLastAccessedGreaterThan),
	}
}

func (s *AutoDeleteCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoDeleteCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoDeleteCondition(input string) (*AutoDeleteCondition, error) {
	vals := map[string]AutoDeleteCondition{
		"createdgreaterthan":      AutoDeleteConditionCreatedGreaterThan,
		"lastaccessedgreaterthan": AutoDeleteConditionLastAccessedGreaterThan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoDeleteCondition(input)
	return &out, nil
}

type BlockedTransformers string

const (
	BlockedTransformersCatTargetEncoder  BlockedTransformers = "CatTargetEncoder"
	BlockedTransformersCountVectorizer   BlockedTransformers = "CountVectorizer"
	BlockedTransformersHashOneHotEncoder BlockedTransformers = "HashOneHotEncoder"
	BlockedTransformersLabelEncoder      BlockedTransformers = "LabelEncoder"
	BlockedTransformersNaiveBayes        BlockedTransformers = "NaiveBayes"
	BlockedTransformersOneHotEncoder     BlockedTransformers = "OneHotEncoder"
	BlockedTransformersTextTargetEncoder BlockedTransformers = "TextTargetEncoder"
	BlockedTransformersTfIdf             BlockedTransformers = "TfIdf"
	BlockedTransformersWoETargetEncoder  BlockedTransformers = "WoETargetEncoder"
	BlockedTransformersWordEmbedding     BlockedTransformers = "WordEmbedding"
)

func PossibleValuesForBlockedTransformers() []string {
	return []string{
		string(BlockedTransformersCatTargetEncoder),
		string(BlockedTransformersCountVectorizer),
		string(BlockedTransformersHashOneHotEncoder),
		string(BlockedTransformersLabelEncoder),
		string(BlockedTransformersNaiveBayes),
		string(BlockedTransformersOneHotEncoder),
		string(BlockedTransformersTextTargetEncoder),
		string(BlockedTransformersTfIdf),
		string(BlockedTransformersWoETargetEncoder),
		string(BlockedTransformersWordEmbedding),
	}
}

func (s *BlockedTransformers) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBlockedTransformers(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBlockedTransformers(input string) (*BlockedTransformers, error) {
	vals := map[string]BlockedTransformers{
		"cattargetencoder":  BlockedTransformersCatTargetEncoder,
		"countvectorizer":   BlockedTransformersCountVectorizer,
		"hashonehotencoder": BlockedTransformersHashOneHotEncoder,
		"labelencoder":      BlockedTransformersLabelEncoder,
		"naivebayes":        BlockedTransformersNaiveBayes,
		"onehotencoder":     BlockedTransformersOneHotEncoder,
		"texttargetencoder": BlockedTransformersTextTargetEncoder,
		"tfidf":             BlockedTransformersTfIdf,
		"woetargetencoder":  BlockedTransformersWoETargetEncoder,
		"wordembedding":     BlockedTransformersWordEmbedding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BlockedTransformers(input)
	return &out, nil
}

type CategoricalDataDriftMetric string

const (
	CategoricalDataDriftMetricJensenShannonDistance    CategoricalDataDriftMetric = "JensenShannonDistance"
	CategoricalDataDriftMetricPearsonsChiSquaredTest   CategoricalDataDriftMetric = "PearsonsChiSquaredTest"
	CategoricalDataDriftMetricPopulationStabilityIndex CategoricalDataDriftMetric = "PopulationStabilityIndex"
)

func PossibleValuesForCategoricalDataDriftMetric() []string {
	return []string{
		string(CategoricalDataDriftMetricJensenShannonDistance),
		string(CategoricalDataDriftMetricPearsonsChiSquaredTest),
		string(CategoricalDataDriftMetricPopulationStabilityIndex),
	}
}

func (s *CategoricalDataDriftMetric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCategoricalDataDriftMetric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCategoricalDataDriftMetric(input string) (*CategoricalDataDriftMetric, error) {
	vals := map[string]CategoricalDataDriftMetric{
		"jensenshannondistance":    CategoricalDataDriftMetricJensenShannonDistance,
		"pearsonschisquaredtest":   CategoricalDataDriftMetricPearsonsChiSquaredTest,
		"populationstabilityindex": CategoricalDataDriftMetricPopulationStabilityIndex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CategoricalDataDriftMetric(input)
	return &out, nil
}

type CategoricalDataQualityMetric string

const (
	CategoricalDataQualityMetricDataTypeErrorRate CategoricalDataQualityMetric = "DataTypeErrorRate"
	CategoricalDataQualityMetricNullValueRate     CategoricalDataQualityMetric = "NullValueRate"
	CategoricalDataQualityMetricOutOfBoundsRate   CategoricalDataQualityMetric = "OutOfBoundsRate"
)

func PossibleValuesForCategoricalDataQualityMetric() []string {
	return []string{
		string(CategoricalDataQualityMetricDataTypeErrorRate),
		string(CategoricalDataQualityMetricNullValueRate),
		string(CategoricalDataQualityMetricOutOfBoundsRate),
	}
}

func (s *CategoricalDataQualityMetric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCategoricalDataQualityMetric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCategoricalDataQualityMetric(input string) (*CategoricalDataQualityMetric, error) {
	vals := map[string]CategoricalDataQualityMetric{
		"datatypeerrorrate": CategoricalDataQualityMetricDataTypeErrorRate,
		"nullvaluerate":     CategoricalDataQualityMetricNullValueRate,
		"outofboundsrate":   CategoricalDataQualityMetricOutOfBoundsRate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CategoricalDataQualityMetric(input)
	return &out, nil
}

type CategoricalPredictionDriftMetric string

const (
	CategoricalPredictionDriftMetricJensenShannonDistance    CategoricalPredictionDriftMetric = "JensenShannonDistance"
	CategoricalPredictionDriftMetricPearsonsChiSquaredTest   CategoricalPredictionDriftMetric = "PearsonsChiSquaredTest"
	CategoricalPredictionDriftMetricPopulationStabilityIndex CategoricalPredictionDriftMetric = "PopulationStabilityIndex"
)

func PossibleValuesForCategoricalPredictionDriftMetric() []string {
	return []string{
		string(CategoricalPredictionDriftMetricJensenShannonDistance),
		string(CategoricalPredictionDriftMetricPearsonsChiSquaredTest),
		string(CategoricalPredictionDriftMetricPopulationStabilityIndex),
	}
}

func (s *CategoricalPredictionDriftMetric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCategoricalPredictionDriftMetric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCategoricalPredictionDriftMetric(input string) (*CategoricalPredictionDriftMetric, error) {
	vals := map[string]CategoricalPredictionDriftMetric{
		"jensenshannondistance":    CategoricalPredictionDriftMetricJensenShannonDistance,
		"pearsonschisquaredtest":   CategoricalPredictionDriftMetricPearsonsChiSquaredTest,
		"populationstabilityindex": CategoricalPredictionDriftMetricPopulationStabilityIndex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CategoricalPredictionDriftMetric(input)
	return &out, nil
}

type ClassificationModelPerformanceMetric string

const (
	ClassificationModelPerformanceMetricAccuracy  ClassificationModelPerformanceMetric = "Accuracy"
	ClassificationModelPerformanceMetricFOneScore ClassificationModelPerformanceMetric = "F1Score"
	ClassificationModelPerformanceMetricPrecision ClassificationModelPerformanceMetric = "Precision"
	ClassificationModelPerformanceMetricRecall    ClassificationModelPerformanceMetric = "Recall"
)

func PossibleValuesForClassificationModelPerformanceMetric() []string {
	return []string{
		string(ClassificationModelPerformanceMetricAccuracy),
		string(ClassificationModelPerformanceMetricFOneScore),
		string(ClassificationModelPerformanceMetricPrecision),
		string(ClassificationModelPerformanceMetricRecall),
	}
}

func (s *ClassificationModelPerformanceMetric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClassificationModelPerformanceMetric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClassificationModelPerformanceMetric(input string) (*ClassificationModelPerformanceMetric, error) {
	vals := map[string]ClassificationModelPerformanceMetric{
		"accuracy":  ClassificationModelPerformanceMetricAccuracy,
		"f1score":   ClassificationModelPerformanceMetricFOneScore,
		"precision": ClassificationModelPerformanceMetricPrecision,
		"recall":    ClassificationModelPerformanceMetricRecall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClassificationModelPerformanceMetric(input)
	return &out, nil
}

type ClassificationModels string

const (
	ClassificationModelsBernoulliNaiveBayes   ClassificationModels = "BernoulliNaiveBayes"
	ClassificationModelsDecisionTree          ClassificationModels = "DecisionTree"
	ClassificationModelsExtremeRandomTrees    ClassificationModels = "ExtremeRandomTrees"
	ClassificationModelsGradientBoosting      ClassificationModels = "GradientBoosting"
	ClassificationModelsKNN                   ClassificationModels = "KNN"
	ClassificationModelsLightGBM              ClassificationModels = "LightGBM"
	ClassificationModelsLinearSVM             ClassificationModels = "LinearSVM"
	ClassificationModelsLogisticRegression    ClassificationModels = "LogisticRegression"
	ClassificationModelsMultinomialNaiveBayes ClassificationModels = "MultinomialNaiveBayes"
	ClassificationModelsRandomForest          ClassificationModels = "RandomForest"
	ClassificationModelsSGD                   ClassificationModels = "SGD"
	ClassificationModelsSVM                   ClassificationModels = "SVM"
	ClassificationModelsXGBoostClassifier     ClassificationModels = "XGBoostClassifier"
)

func PossibleValuesForClassificationModels() []string {
	return []string{
		string(ClassificationModelsBernoulliNaiveBayes),
		string(ClassificationModelsDecisionTree),
		string(ClassificationModelsExtremeRandomTrees),
		string(ClassificationModelsGradientBoosting),
		string(ClassificationModelsKNN),
		string(ClassificationModelsLightGBM),
		string(ClassificationModelsLinearSVM),
		string(ClassificationModelsLogisticRegression),
		string(ClassificationModelsMultinomialNaiveBayes),
		string(ClassificationModelsRandomForest),
		string(ClassificationModelsSGD),
		string(ClassificationModelsSVM),
		string(ClassificationModelsXGBoostClassifier),
	}
}

func (s *ClassificationModels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClassificationModels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClassificationModels(input string) (*ClassificationModels, error) {
	vals := map[string]ClassificationModels{
		"bernoullinaivebayes":   ClassificationModelsBernoulliNaiveBayes,
		"decisiontree":          ClassificationModelsDecisionTree,
		"extremerandomtrees":    ClassificationModelsExtremeRandomTrees,
		"gradientboosting":      ClassificationModelsGradientBoosting,
		"knn":                   ClassificationModelsKNN,
		"lightgbm":              ClassificationModelsLightGBM,
		"linearsvm":             ClassificationModelsLinearSVM,
		"logisticregression":    ClassificationModelsLogisticRegression,
		"multinomialnaivebayes": ClassificationModelsMultinomialNaiveBayes,
		"randomforest":          ClassificationModelsRandomForest,
		"sgd":                   ClassificationModelsSGD,
		"svm":                   ClassificationModelsSVM,
		"xgboostclassifier":     ClassificationModelsXGBoostClassifier,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClassificationModels(input)
	return &out, nil
}

type ClassificationMultilabelPrimaryMetrics string

const (
	ClassificationMultilabelPrimaryMetricsAUCWeighted                   ClassificationMultilabelPrimaryMetrics = "AUCWeighted"
	ClassificationMultilabelPrimaryMetricsAccuracy                      ClassificationMultilabelPrimaryMetrics = "Accuracy"
	ClassificationMultilabelPrimaryMetricsAveragePrecisionScoreWeighted ClassificationMultilabelPrimaryMetrics = "AveragePrecisionScoreWeighted"
	ClassificationMultilabelPrimaryMetricsIOU                           ClassificationMultilabelPrimaryMetrics = "IOU"
	ClassificationMultilabelPrimaryMetricsNormMacroRecall               ClassificationMultilabelPrimaryMetrics = "NormMacroRecall"
	ClassificationMultilabelPrimaryMetricsPrecisionScoreWeighted        ClassificationMultilabelPrimaryMetrics = "PrecisionScoreWeighted"
)

func PossibleValuesForClassificationMultilabelPrimaryMetrics() []string {
	return []string{
		string(ClassificationMultilabelPrimaryMetricsAUCWeighted),
		string(ClassificationMultilabelPrimaryMetricsAccuracy),
		string(ClassificationMultilabelPrimaryMetricsAveragePrecisionScoreWeighted),
		string(ClassificationMultilabelPrimaryMetricsIOU),
		string(ClassificationMultilabelPrimaryMetricsNormMacroRecall),
		string(ClassificationMultilabelPrimaryMetricsPrecisionScoreWeighted),
	}
}

func (s *ClassificationMultilabelPrimaryMetrics) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClassificationMultilabelPrimaryMetrics(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClassificationMultilabelPrimaryMetrics(input string) (*ClassificationMultilabelPrimaryMetrics, error) {
	vals := map[string]ClassificationMultilabelPrimaryMetrics{
		"aucweighted":                   ClassificationMultilabelPrimaryMetricsAUCWeighted,
		"accuracy":                      ClassificationMultilabelPrimaryMetricsAccuracy,
		"averageprecisionscoreweighted": ClassificationMultilabelPrimaryMetricsAveragePrecisionScoreWeighted,
		"iou":                           ClassificationMultilabelPrimaryMetricsIOU,
		"normmacrorecall":               ClassificationMultilabelPrimaryMetricsNormMacroRecall,
		"precisionscoreweighted":        ClassificationMultilabelPrimaryMetricsPrecisionScoreWeighted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClassificationMultilabelPrimaryMetrics(input)
	return &out, nil
}

type ClassificationPrimaryMetrics string

const (
	ClassificationPrimaryMetricsAUCWeighted                   ClassificationPrimaryMetrics = "AUCWeighted"
	ClassificationPrimaryMetricsAccuracy                      ClassificationPrimaryMetrics = "Accuracy"
	ClassificationPrimaryMetricsAveragePrecisionScoreWeighted ClassificationPrimaryMetrics = "AveragePrecisionScoreWeighted"
	ClassificationPrimaryMetricsNormMacroRecall               ClassificationPrimaryMetrics = "NormMacroRecall"
	ClassificationPrimaryMetricsPrecisionScoreWeighted        ClassificationPrimaryMetrics = "PrecisionScoreWeighted"
)

func PossibleValuesForClassificationPrimaryMetrics() []string {
	return []string{
		string(ClassificationPrimaryMetricsAUCWeighted),
		string(ClassificationPrimaryMetricsAccuracy),
		string(ClassificationPrimaryMetricsAveragePrecisionScoreWeighted),
		string(ClassificationPrimaryMetricsNormMacroRecall),
		string(ClassificationPrimaryMetricsPrecisionScoreWeighted),
	}
}

func (s *ClassificationPrimaryMetrics) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClassificationPrimaryMetrics(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClassificationPrimaryMetrics(input string) (*ClassificationPrimaryMetrics, error) {
	vals := map[string]ClassificationPrimaryMetrics{
		"aucweighted":                   ClassificationPrimaryMetricsAUCWeighted,
		"accuracy":                      ClassificationPrimaryMetricsAccuracy,
		"averageprecisionscoreweighted": ClassificationPrimaryMetricsAveragePrecisionScoreWeighted,
		"normmacrorecall":               ClassificationPrimaryMetricsNormMacroRecall,
		"precisionscoreweighted":        ClassificationPrimaryMetricsPrecisionScoreWeighted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClassificationPrimaryMetrics(input)
	return &out, nil
}

type DataImportSourceType string

const (
	DataImportSourceTypeDatabase   DataImportSourceType = "database"
	DataImportSourceTypeFileSystem DataImportSourceType = "file_system"
)

func PossibleValuesForDataImportSourceType() []string {
	return []string{
		string(DataImportSourceTypeDatabase),
		string(DataImportSourceTypeFileSystem),
	}
}

func (s *DataImportSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataImportSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataImportSourceType(input string) (*DataImportSourceType, error) {
	vals := map[string]DataImportSourceType{
		"database":    DataImportSourceTypeDatabase,
		"file_system": DataImportSourceTypeFileSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataImportSourceType(input)
	return &out, nil
}

type DataType string

const (
	DataTypeMltable   DataType = "mltable"
	DataTypeUriFile   DataType = "uri_file"
	DataTypeUriFolder DataType = "uri_folder"
)

func PossibleValuesForDataType() []string {
	return []string{
		string(DataTypeMltable),
		string(DataTypeUriFile),
		string(DataTypeUriFolder),
	}
}

func (s *DataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataType(input string) (*DataType, error) {
	vals := map[string]DataType{
		"mltable":    DataTypeMltable,
		"uri_file":   DataTypeUriFile,
		"uri_folder": DataTypeUriFolder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataType(input)
	return &out, nil
}

type DistributionType string

const (
	DistributionTypeMpi        DistributionType = "Mpi"
	DistributionTypePyTorch    DistributionType = "PyTorch"
	DistributionTypeRay        DistributionType = "Ray"
	DistributionTypeTensorFlow DistributionType = "TensorFlow"
)

func PossibleValuesForDistributionType() []string {
	return []string{
		string(DistributionTypeMpi),
		string(DistributionTypePyTorch),
		string(DistributionTypeRay),
		string(DistributionTypeTensorFlow),
	}
}

func (s *DistributionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDistributionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDistributionType(input string) (*DistributionType, error) {
	vals := map[string]DistributionType{
		"mpi":        DistributionTypeMpi,
		"pytorch":    DistributionTypePyTorch,
		"ray":        DistributionTypeRay,
		"tensorflow": DistributionTypeTensorFlow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DistributionType(input)
	return &out, nil
}

type EarlyTerminationPolicyType string

const (
	EarlyTerminationPolicyTypeBandit              EarlyTerminationPolicyType = "Bandit"
	EarlyTerminationPolicyTypeMedianStopping      EarlyTerminationPolicyType = "MedianStopping"
	EarlyTerminationPolicyTypeTruncationSelection EarlyTerminationPolicyType = "TruncationSelection"
)

func PossibleValuesForEarlyTerminationPolicyType() []string {
	return []string{
		string(EarlyTerminationPolicyTypeBandit),
		string(EarlyTerminationPolicyTypeMedianStopping),
		string(EarlyTerminationPolicyTypeTruncationSelection),
	}
}

func (s *EarlyTerminationPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEarlyTerminationPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEarlyTerminationPolicyType(input string) (*EarlyTerminationPolicyType, error) {
	vals := map[string]EarlyTerminationPolicyType{
		"bandit":              EarlyTerminationPolicyTypeBandit,
		"medianstopping":      EarlyTerminationPolicyTypeMedianStopping,
		"truncationselection": EarlyTerminationPolicyTypeTruncationSelection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EarlyTerminationPolicyType(input)
	return &out, nil
}

type EmailNotificationEnableType string

const (
	EmailNotificationEnableTypeJobCancelled EmailNotificationEnableType = "JobCancelled"
	EmailNotificationEnableTypeJobCompleted EmailNotificationEnableType = "JobCompleted"
	EmailNotificationEnableTypeJobFailed    EmailNotificationEnableType = "JobFailed"
)

func PossibleValuesForEmailNotificationEnableType() []string {
	return []string{
		string(EmailNotificationEnableTypeJobCancelled),
		string(EmailNotificationEnableTypeJobCompleted),
		string(EmailNotificationEnableTypeJobFailed),
	}
}

func (s *EmailNotificationEnableType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailNotificationEnableType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailNotificationEnableType(input string) (*EmailNotificationEnableType, error) {
	vals := map[string]EmailNotificationEnableType{
		"jobcancelled": EmailNotificationEnableTypeJobCancelled,
		"jobcompleted": EmailNotificationEnableTypeJobCompleted,
		"jobfailed":    EmailNotificationEnableTypeJobFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailNotificationEnableType(input)
	return &out, nil
}

type FeatureAttributionMetric string

const (
	FeatureAttributionMetricNormalizedDiscountedCumulativeGain FeatureAttributionMetric = "NormalizedDiscountedCumulativeGain"
)

func PossibleValuesForFeatureAttributionMetric() []string {
	return []string{
		string(FeatureAttributionMetricNormalizedDiscountedCumulativeGain),
	}
}

func (s *FeatureAttributionMetric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureAttributionMetric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureAttributionMetric(input string) (*FeatureAttributionMetric, error) {
	vals := map[string]FeatureAttributionMetric{
		"normalizeddiscountedcumulativegain": FeatureAttributionMetricNormalizedDiscountedCumulativeGain,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureAttributionMetric(input)
	return &out, nil
}

type FeatureLags string

const (
	FeatureLagsAuto FeatureLags = "Auto"
	FeatureLagsNone FeatureLags = "None"
)

func PossibleValuesForFeatureLags() []string {
	return []string{
		string(FeatureLagsAuto),
		string(FeatureLagsNone),
	}
}

func (s *FeatureLags) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureLags(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureLags(input string) (*FeatureLags, error) {
	vals := map[string]FeatureLags{
		"auto": FeatureLagsAuto,
		"none": FeatureLagsNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureLags(input)
	return &out, nil
}

type FeaturizationMode string

const (
	FeaturizationModeAuto   FeaturizationMode = "Auto"
	FeaturizationModeCustom FeaturizationMode = "Custom"
	FeaturizationModeOff    FeaturizationMode = "Off"
)

func PossibleValuesForFeaturizationMode() []string {
	return []string{
		string(FeaturizationModeAuto),
		string(FeaturizationModeCustom),
		string(FeaturizationModeOff),
	}
}

func (s *FeaturizationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeaturizationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeaturizationMode(input string) (*FeaturizationMode, error) {
	vals := map[string]FeaturizationMode{
		"auto":   FeaturizationModeAuto,
		"custom": FeaturizationModeCustom,
		"off":    FeaturizationModeOff,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeaturizationMode(input)
	return &out, nil
}

type ForecastHorizonMode string

const (
	ForecastHorizonModeAuto   ForecastHorizonMode = "Auto"
	ForecastHorizonModeCustom ForecastHorizonMode = "Custom"
)

func PossibleValuesForForecastHorizonMode() []string {
	return []string{
		string(ForecastHorizonModeAuto),
		string(ForecastHorizonModeCustom),
	}
}

func (s *ForecastHorizonMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseForecastHorizonMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseForecastHorizonMode(input string) (*ForecastHorizonMode, error) {
	vals := map[string]ForecastHorizonMode{
		"auto":   ForecastHorizonModeAuto,
		"custom": ForecastHorizonModeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForecastHorizonMode(input)
	return &out, nil
}

type ForecastingModels string

const (
	ForecastingModelsArimax               ForecastingModels = "Arimax"
	ForecastingModelsAutoArima            ForecastingModels = "AutoArima"
	ForecastingModelsAverage              ForecastingModels = "Average"
	ForecastingModelsDecisionTree         ForecastingModels = "DecisionTree"
	ForecastingModelsElasticNet           ForecastingModels = "ElasticNet"
	ForecastingModelsExponentialSmoothing ForecastingModels = "ExponentialSmoothing"
	ForecastingModelsExtremeRandomTrees   ForecastingModels = "ExtremeRandomTrees"
	ForecastingModelsGradientBoosting     ForecastingModels = "GradientBoosting"
	ForecastingModelsKNN                  ForecastingModels = "KNN"
	ForecastingModelsLassoLars            ForecastingModels = "LassoLars"
	ForecastingModelsLightGBM             ForecastingModels = "LightGBM"
	ForecastingModelsNaive                ForecastingModels = "Naive"
	ForecastingModelsProphet              ForecastingModels = "Prophet"
	ForecastingModelsRandomForest         ForecastingModels = "RandomForest"
	ForecastingModelsSGD                  ForecastingModels = "SGD"
	ForecastingModelsSeasonalAverage      ForecastingModels = "SeasonalAverage"
	ForecastingModelsSeasonalNaive        ForecastingModels = "SeasonalNaive"
	ForecastingModelsTCNForecaster        ForecastingModels = "TCNForecaster"
	ForecastingModelsXGBoostRegressor     ForecastingModels = "XGBoostRegressor"
)

func PossibleValuesForForecastingModels() []string {
	return []string{
		string(ForecastingModelsArimax),
		string(ForecastingModelsAutoArima),
		string(ForecastingModelsAverage),
		string(ForecastingModelsDecisionTree),
		string(ForecastingModelsElasticNet),
		string(ForecastingModelsExponentialSmoothing),
		string(ForecastingModelsExtremeRandomTrees),
		string(ForecastingModelsGradientBoosting),
		string(ForecastingModelsKNN),
		string(ForecastingModelsLassoLars),
		string(ForecastingModelsLightGBM),
		string(ForecastingModelsNaive),
		string(ForecastingModelsProphet),
		string(ForecastingModelsRandomForest),
		string(ForecastingModelsSGD),
		string(ForecastingModelsSeasonalAverage),
		string(ForecastingModelsSeasonalNaive),
		string(ForecastingModelsTCNForecaster),
		string(ForecastingModelsXGBoostRegressor),
	}
}

func (s *ForecastingModels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseForecastingModels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseForecastingModels(input string) (*ForecastingModels, error) {
	vals := map[string]ForecastingModels{
		"arimax":               ForecastingModelsArimax,
		"autoarima":            ForecastingModelsAutoArima,
		"average":              ForecastingModelsAverage,
		"decisiontree":         ForecastingModelsDecisionTree,
		"elasticnet":           ForecastingModelsElasticNet,
		"exponentialsmoothing": ForecastingModelsExponentialSmoothing,
		"extremerandomtrees":   ForecastingModelsExtremeRandomTrees,
		"gradientboosting":     ForecastingModelsGradientBoosting,
		"knn":                  ForecastingModelsKNN,
		"lassolars":            ForecastingModelsLassoLars,
		"lightgbm":             ForecastingModelsLightGBM,
		"naive":                ForecastingModelsNaive,
		"prophet":              ForecastingModelsProphet,
		"randomforest":         ForecastingModelsRandomForest,
		"sgd":                  ForecastingModelsSGD,
		"seasonalaverage":      ForecastingModelsSeasonalAverage,
		"seasonalnaive":        ForecastingModelsSeasonalNaive,
		"tcnforecaster":        ForecastingModelsTCNForecaster,
		"xgboostregressor":     ForecastingModelsXGBoostRegressor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForecastingModels(input)
	return &out, nil
}

type ForecastingPrimaryMetrics string

const (
	ForecastingPrimaryMetricsNormalizedMeanAbsoluteError    ForecastingPrimaryMetrics = "NormalizedMeanAbsoluteError"
	ForecastingPrimaryMetricsNormalizedRootMeanSquaredError ForecastingPrimaryMetrics = "NormalizedRootMeanSquaredError"
	ForecastingPrimaryMetricsRTwoScore                      ForecastingPrimaryMetrics = "R2Score"
	ForecastingPrimaryMetricsSpearmanCorrelation            ForecastingPrimaryMetrics = "SpearmanCorrelation"
)

func PossibleValuesForForecastingPrimaryMetrics() []string {
	return []string{
		string(ForecastingPrimaryMetricsNormalizedMeanAbsoluteError),
		string(ForecastingPrimaryMetricsNormalizedRootMeanSquaredError),
		string(ForecastingPrimaryMetricsRTwoScore),
		string(ForecastingPrimaryMetricsSpearmanCorrelation),
	}
}

func (s *ForecastingPrimaryMetrics) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseForecastingPrimaryMetrics(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseForecastingPrimaryMetrics(input string) (*ForecastingPrimaryMetrics, error) {
	vals := map[string]ForecastingPrimaryMetrics{
		"normalizedmeanabsoluteerror":    ForecastingPrimaryMetricsNormalizedMeanAbsoluteError,
		"normalizedrootmeansquarederror": ForecastingPrimaryMetricsNormalizedRootMeanSquaredError,
		"r2score":                        ForecastingPrimaryMetricsRTwoScore,
		"spearmancorrelation":            ForecastingPrimaryMetricsSpearmanCorrelation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForecastingPrimaryMetrics(input)
	return &out, nil
}

type Goal string

const (
	GoalMaximize Goal = "Maximize"
	GoalMinimize Goal = "Minimize"
)

func PossibleValuesForGoal() []string {
	return []string{
		string(GoalMaximize),
		string(GoalMinimize),
	}
}

func (s *Goal) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGoal(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGoal(input string) (*Goal, error) {
	vals := map[string]Goal{
		"maximize": GoalMaximize,
		"minimize": GoalMinimize,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Goal(input)
	return &out, nil
}

type IdentityConfigurationType string

const (
	IdentityConfigurationTypeAMLToken     IdentityConfigurationType = "AMLToken"
	IdentityConfigurationTypeManaged      IdentityConfigurationType = "Managed"
	IdentityConfigurationTypeUserIdentity IdentityConfigurationType = "UserIdentity"
)

func PossibleValuesForIdentityConfigurationType() []string {
	return []string{
		string(IdentityConfigurationTypeAMLToken),
		string(IdentityConfigurationTypeManaged),
		string(IdentityConfigurationTypeUserIdentity),
	}
}

func (s *IdentityConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityConfigurationType(input string) (*IdentityConfigurationType, error) {
	vals := map[string]IdentityConfigurationType{
		"amltoken":     IdentityConfigurationTypeAMLToken,
		"managed":      IdentityConfigurationTypeManaged,
		"useridentity": IdentityConfigurationTypeUserIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityConfigurationType(input)
	return &out, nil
}

type ImageAnnotationType string

const (
	ImageAnnotationTypeBoundingBox          ImageAnnotationType = "BoundingBox"
	ImageAnnotationTypeClassification       ImageAnnotationType = "Classification"
	ImageAnnotationTypeInstanceSegmentation ImageAnnotationType = "InstanceSegmentation"
)

func PossibleValuesForImageAnnotationType() []string {
	return []string{
		string(ImageAnnotationTypeBoundingBox),
		string(ImageAnnotationTypeClassification),
		string(ImageAnnotationTypeInstanceSegmentation),
	}
}

func (s *ImageAnnotationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImageAnnotationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImageAnnotationType(input string) (*ImageAnnotationType, error) {
	vals := map[string]ImageAnnotationType{
		"boundingbox":          ImageAnnotationTypeBoundingBox,
		"classification":       ImageAnnotationTypeClassification,
		"instancesegmentation": ImageAnnotationTypeInstanceSegmentation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImageAnnotationType(input)
	return &out, nil
}

type IncrementalDataRefresh string

const (
	IncrementalDataRefreshDisabled IncrementalDataRefresh = "Disabled"
	IncrementalDataRefreshEnabled  IncrementalDataRefresh = "Enabled"
)

func PossibleValuesForIncrementalDataRefresh() []string {
	return []string{
		string(IncrementalDataRefreshDisabled),
		string(IncrementalDataRefreshEnabled),
	}
}

func (s *IncrementalDataRefresh) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncrementalDataRefresh(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncrementalDataRefresh(input string) (*IncrementalDataRefresh, error) {
	vals := map[string]IncrementalDataRefresh{
		"disabled": IncrementalDataRefreshDisabled,
		"enabled":  IncrementalDataRefreshEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncrementalDataRefresh(input)
	return &out, nil
}

type InputDeliveryMode string

const (
	InputDeliveryModeDirect         InputDeliveryMode = "Direct"
	InputDeliveryModeDownload       InputDeliveryMode = "Download"
	InputDeliveryModeEvalDownload   InputDeliveryMode = "EvalDownload"
	InputDeliveryModeEvalMount      InputDeliveryMode = "EvalMount"
	InputDeliveryModeReadOnlyMount  InputDeliveryMode = "ReadOnlyMount"
	InputDeliveryModeReadWriteMount InputDeliveryMode = "ReadWriteMount"
)

func PossibleValuesForInputDeliveryMode() []string {
	return []string{
		string(InputDeliveryModeDirect),
		string(InputDeliveryModeDownload),
		string(InputDeliveryModeEvalDownload),
		string(InputDeliveryModeEvalMount),
		string(InputDeliveryModeReadOnlyMount),
		string(InputDeliveryModeReadWriteMount),
	}
}

func (s *InputDeliveryMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInputDeliveryMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInputDeliveryMode(input string) (*InputDeliveryMode, error) {
	vals := map[string]InputDeliveryMode{
		"direct":         InputDeliveryModeDirect,
		"download":       InputDeliveryModeDownload,
		"evaldownload":   InputDeliveryModeEvalDownload,
		"evalmount":      InputDeliveryModeEvalMount,
		"readonlymount":  InputDeliveryModeReadOnlyMount,
		"readwritemount": InputDeliveryModeReadWriteMount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InputDeliveryMode(input)
	return &out, nil
}

type InstanceSegmentationPrimaryMetrics string

const (
	InstanceSegmentationPrimaryMetricsMeanAveragePrecision InstanceSegmentationPrimaryMetrics = "MeanAveragePrecision"
)

func PossibleValuesForInstanceSegmentationPrimaryMetrics() []string {
	return []string{
		string(InstanceSegmentationPrimaryMetricsMeanAveragePrecision),
	}
}

func (s *InstanceSegmentationPrimaryMetrics) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInstanceSegmentationPrimaryMetrics(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInstanceSegmentationPrimaryMetrics(input string) (*InstanceSegmentationPrimaryMetrics, error) {
	vals := map[string]InstanceSegmentationPrimaryMetrics{
		"meanaverageprecision": InstanceSegmentationPrimaryMetricsMeanAveragePrecision,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstanceSegmentationPrimaryMetrics(input)
	return &out, nil
}

type JobInputType string

const (
	JobInputTypeCustomModel JobInputType = "custom_model"
	JobInputTypeLiteral     JobInputType = "literal"
	JobInputTypeMlflowModel JobInputType = "mlflow_model"
	JobInputTypeMltable     JobInputType = "mltable"
	JobInputTypeTritonModel JobInputType = "triton_model"
	JobInputTypeUriFile     JobInputType = "uri_file"
	JobInputTypeUriFolder   JobInputType = "uri_folder"
)

func PossibleValuesForJobInputType() []string {
	return []string{
		string(JobInputTypeCustomModel),
		string(JobInputTypeLiteral),
		string(JobInputTypeMlflowModel),
		string(JobInputTypeMltable),
		string(JobInputTypeTritonModel),
		string(JobInputTypeUriFile),
		string(JobInputTypeUriFolder),
	}
}

func (s *JobInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobInputType(input string) (*JobInputType, error) {
	vals := map[string]JobInputType{
		"custom_model": JobInputTypeCustomModel,
		"literal":      JobInputTypeLiteral,
		"mlflow_model": JobInputTypeMlflowModel,
		"mltable":      JobInputTypeMltable,
		"triton_model": JobInputTypeTritonModel,
		"uri_file":     JobInputTypeUriFile,
		"uri_folder":   JobInputTypeUriFolder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobInputType(input)
	return &out, nil
}

type JobLimitsType string

const (
	JobLimitsTypeCommand JobLimitsType = "Command"
	JobLimitsTypeSweep   JobLimitsType = "Sweep"
)

func PossibleValuesForJobLimitsType() []string {
	return []string{
		string(JobLimitsTypeCommand),
		string(JobLimitsTypeSweep),
	}
}

func (s *JobLimitsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobLimitsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobLimitsType(input string) (*JobLimitsType, error) {
	vals := map[string]JobLimitsType{
		"command": JobLimitsTypeCommand,
		"sweep":   JobLimitsTypeSweep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobLimitsType(input)
	return &out, nil
}

type JobOutputType string

const (
	JobOutputTypeCustomModel JobOutputType = "custom_model"
	JobOutputTypeMlflowModel JobOutputType = "mlflow_model"
	JobOutputTypeMltable     JobOutputType = "mltable"
	JobOutputTypeTritonModel JobOutputType = "triton_model"
	JobOutputTypeUriFile     JobOutputType = "uri_file"
	JobOutputTypeUriFolder   JobOutputType = "uri_folder"
)

func PossibleValuesForJobOutputType() []string {
	return []string{
		string(JobOutputTypeCustomModel),
		string(JobOutputTypeMlflowModel),
		string(JobOutputTypeMltable),
		string(JobOutputTypeTritonModel),
		string(JobOutputTypeUriFile),
		string(JobOutputTypeUriFolder),
	}
}

func (s *JobOutputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobOutputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobOutputType(input string) (*JobOutputType, error) {
	vals := map[string]JobOutputType{
		"custom_model": JobOutputTypeCustomModel,
		"mlflow_model": JobOutputTypeMlflowModel,
		"mltable":      JobOutputTypeMltable,
		"triton_model": JobOutputTypeTritonModel,
		"uri_file":     JobOutputTypeUriFile,
		"uri_folder":   JobOutputTypeUriFolder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobOutputType(input)
	return &out, nil
}

type JobProvisioningState string

const (
	JobProvisioningStateCanceled   JobProvisioningState = "Canceled"
	JobProvisioningStateFailed     JobProvisioningState = "Failed"
	JobProvisioningStateInProgress JobProvisioningState = "InProgress"
	JobProvisioningStateSucceeded  JobProvisioningState = "Succeeded"
)

func PossibleValuesForJobProvisioningState() []string {
	return []string{
		string(JobProvisioningStateCanceled),
		string(JobProvisioningStateFailed),
		string(JobProvisioningStateInProgress),
		string(JobProvisioningStateSucceeded),
	}
}

func (s *JobProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobProvisioningState(input string) (*JobProvisioningState, error) {
	vals := map[string]JobProvisioningState{
		"canceled":   JobProvisioningStateCanceled,
		"failed":     JobProvisioningStateFailed,
		"inprogress": JobProvisioningStateInProgress,
		"succeeded":  JobProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobProvisioningState(input)
	return &out, nil
}

type JobStatus string

const (
	JobStatusCancelRequested JobStatus = "CancelRequested"
	JobStatusCanceled        JobStatus = "Canceled"
	JobStatusCompleted       JobStatus = "Completed"
	JobStatusFailed          JobStatus = "Failed"
	JobStatusFinalizing      JobStatus = "Finalizing"
	JobStatusNotResponding   JobStatus = "NotResponding"
	JobStatusNotStarted      JobStatus = "NotStarted"
	JobStatusPaused          JobStatus = "Paused"
	JobStatusPreparing       JobStatus = "Preparing"
	JobStatusProvisioning    JobStatus = "Provisioning"
	JobStatusQueued          JobStatus = "Queued"
	JobStatusRunning         JobStatus = "Running"
	JobStatusScheduled       JobStatus = "Scheduled"
	JobStatusStarting        JobStatus = "Starting"
	JobStatusUnknown         JobStatus = "Unknown"
)

func PossibleValuesForJobStatus() []string {
	return []string{
		string(JobStatusCancelRequested),
		string(JobStatusCanceled),
		string(JobStatusCompleted),
		string(JobStatusFailed),
		string(JobStatusFinalizing),
		string(JobStatusNotResponding),
		string(JobStatusNotStarted),
		string(JobStatusPaused),
		string(JobStatusPreparing),
		string(JobStatusProvisioning),
		string(JobStatusQueued),
		string(JobStatusRunning),
		string(JobStatusScheduled),
		string(JobStatusStarting),
		string(JobStatusUnknown),
	}
}

func (s *JobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobStatus(input string) (*JobStatus, error) {
	vals := map[string]JobStatus{
		"cancelrequested": JobStatusCancelRequested,
		"canceled":        JobStatusCanceled,
		"completed":       JobStatusCompleted,
		"failed":          JobStatusFailed,
		"finalizing":      JobStatusFinalizing,
		"notresponding":   JobStatusNotResponding,
		"notstarted":      JobStatusNotStarted,
		"paused":          JobStatusPaused,
		"preparing":       JobStatusPreparing,
		"provisioning":    JobStatusProvisioning,
		"queued":          JobStatusQueued,
		"running":         JobStatusRunning,
		"scheduled":       JobStatusScheduled,
		"starting":        JobStatusStarting,
		"unknown":         JobStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobStatus(input)
	return &out, nil
}

type JobTier string

const (
	JobTierBasic    JobTier = "Basic"
	JobTierPremium  JobTier = "Premium"
	JobTierSpot     JobTier = "Spot"
	JobTierStandard JobTier = "Standard"
)

func PossibleValuesForJobTier() []string {
	return []string{
		string(JobTierBasic),
		string(JobTierPremium),
		string(JobTierSpot),
		string(JobTierStandard),
	}
}

func (s *JobTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobTier(input string) (*JobTier, error) {
	vals := map[string]JobTier{
		"basic":    JobTierBasic,
		"premium":  JobTierPremium,
		"spot":     JobTierSpot,
		"standard": JobTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobTier(input)
	return &out, nil
}

type JobType string

const (
	JobTypeAutoML   JobType = "AutoML"
	JobTypeCommand  JobType = "Command"
	JobTypeLabeling JobType = "Labeling"
	JobTypePipeline JobType = "Pipeline"
	JobTypeSpark    JobType = "Spark"
	JobTypeSweep    JobType = "Sweep"
)

func PossibleValuesForJobType() []string {
	return []string{
		string(JobTypeAutoML),
		string(JobTypeCommand),
		string(JobTypeLabeling),
		string(JobTypePipeline),
		string(JobTypeSpark),
		string(JobTypeSweep),
	}
}

func (s *JobType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobType(input string) (*JobType, error) {
	vals := map[string]JobType{
		"automl":   JobTypeAutoML,
		"command":  JobTypeCommand,
		"labeling": JobTypeLabeling,
		"pipeline": JobTypePipeline,
		"spark":    JobTypeSpark,
		"sweep":    JobTypeSweep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobType(input)
	return &out, nil
}

type LearningRateScheduler string

const (
	LearningRateSchedulerNone         LearningRateScheduler = "None"
	LearningRateSchedulerStep         LearningRateScheduler = "Step"
	LearningRateSchedulerWarmupCosine LearningRateScheduler = "WarmupCosine"
)

func PossibleValuesForLearningRateScheduler() []string {
	return []string{
		string(LearningRateSchedulerNone),
		string(LearningRateSchedulerStep),
		string(LearningRateSchedulerWarmupCosine),
	}
}

func (s *LearningRateScheduler) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLearningRateScheduler(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLearningRateScheduler(input string) (*LearningRateScheduler, error) {
	vals := map[string]LearningRateScheduler{
		"none":         LearningRateSchedulerNone,
		"step":         LearningRateSchedulerStep,
		"warmupcosine": LearningRateSchedulerWarmupCosine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LearningRateScheduler(input)
	return &out, nil
}

type LogTrainingMetrics string

const (
	LogTrainingMetricsDisable LogTrainingMetrics = "Disable"
	LogTrainingMetricsEnable  LogTrainingMetrics = "Enable"
)

func PossibleValuesForLogTrainingMetrics() []string {
	return []string{
		string(LogTrainingMetricsDisable),
		string(LogTrainingMetricsEnable),
	}
}

func (s *LogTrainingMetrics) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogTrainingMetrics(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogTrainingMetrics(input string) (*LogTrainingMetrics, error) {
	vals := map[string]LogTrainingMetrics{
		"disable": LogTrainingMetricsDisable,
		"enable":  LogTrainingMetricsEnable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogTrainingMetrics(input)
	return &out, nil
}

type LogValidationLoss string

const (
	LogValidationLossDisable LogValidationLoss = "Disable"
	LogValidationLossEnable  LogValidationLoss = "Enable"
)

func PossibleValuesForLogValidationLoss() []string {
	return []string{
		string(LogValidationLossDisable),
		string(LogValidationLossEnable),
	}
}

func (s *LogValidationLoss) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogValidationLoss(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogValidationLoss(input string) (*LogValidationLoss, error) {
	vals := map[string]LogValidationLoss{
		"disable": LogValidationLossDisable,
		"enable":  LogValidationLossEnable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogValidationLoss(input)
	return &out, nil
}

type LogVerbosity string

const (
	LogVerbosityCritical LogVerbosity = "Critical"
	LogVerbosityDebug    LogVerbosity = "Debug"
	LogVerbosityError    LogVerbosity = "Error"
	LogVerbosityInfo     LogVerbosity = "Info"
	LogVerbosityNotSet   LogVerbosity = "NotSet"
	LogVerbosityWarning  LogVerbosity = "Warning"
)

func PossibleValuesForLogVerbosity() []string {
	return []string{
		string(LogVerbosityCritical),
		string(LogVerbosityDebug),
		string(LogVerbosityError),
		string(LogVerbosityInfo),
		string(LogVerbosityNotSet),
		string(LogVerbosityWarning),
	}
}

func (s *LogVerbosity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogVerbosity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogVerbosity(input string) (*LogVerbosity, error) {
	vals := map[string]LogVerbosity{
		"critical": LogVerbosityCritical,
		"debug":    LogVerbosityDebug,
		"error":    LogVerbosityError,
		"info":     LogVerbosityInfo,
		"notset":   LogVerbosityNotSet,
		"warning":  LogVerbosityWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogVerbosity(input)
	return &out, nil
}

type MLAssistConfigurationType string

const (
	MLAssistConfigurationTypeDisabled MLAssistConfigurationType = "Disabled"
	MLAssistConfigurationTypeEnabled  MLAssistConfigurationType = "Enabled"
)

func PossibleValuesForMLAssistConfigurationType() []string {
	return []string{
		string(MLAssistConfigurationTypeDisabled),
		string(MLAssistConfigurationTypeEnabled),
	}
}

func (s *MLAssistConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMLAssistConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMLAssistConfigurationType(input string) (*MLAssistConfigurationType, error) {
	vals := map[string]MLAssistConfigurationType{
		"disabled": MLAssistConfigurationTypeDisabled,
		"enabled":  MLAssistConfigurationTypeEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MLAssistConfigurationType(input)
	return &out, nil
}

type MLFlowAutologgerState string

const (
	MLFlowAutologgerStateDisabled MLFlowAutologgerState = "Disabled"
	MLFlowAutologgerStateEnabled  MLFlowAutologgerState = "Enabled"
)

func PossibleValuesForMLFlowAutologgerState() []string {
	return []string{
		string(MLFlowAutologgerStateDisabled),
		string(MLFlowAutologgerStateEnabled),
	}
}

func (s *MLFlowAutologgerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMLFlowAutologgerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMLFlowAutologgerState(input string) (*MLFlowAutologgerState, error) {
	vals := map[string]MLFlowAutologgerState{
		"disabled": MLFlowAutologgerStateDisabled,
		"enabled":  MLFlowAutologgerStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MLFlowAutologgerState(input)
	return &out, nil
}

type MediaType string

const (
	MediaTypeImage MediaType = "Image"
	MediaTypeText  MediaType = "Text"
)

func PossibleValuesForMediaType() []string {
	return []string{
		string(MediaTypeImage),
		string(MediaTypeText),
	}
}

func (s *MediaType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaType(input string) (*MediaType, error) {
	vals := map[string]MediaType{
		"image": MediaTypeImage,
		"text":  MediaTypeText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaType(input)
	return &out, nil
}

type ModelSize string

const (
	ModelSizeExtraLarge ModelSize = "ExtraLarge"
	ModelSizeLarge      ModelSize = "Large"
	ModelSizeMedium     ModelSize = "Medium"
	ModelSizeNone       ModelSize = "None"
	ModelSizeSmall      ModelSize = "Small"
)

func PossibleValuesForModelSize() []string {
	return []string{
		string(ModelSizeExtraLarge),
		string(ModelSizeLarge),
		string(ModelSizeMedium),
		string(ModelSizeNone),
		string(ModelSizeSmall),
	}
}

func (s *ModelSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseModelSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseModelSize(input string) (*ModelSize, error) {
	vals := map[string]ModelSize{
		"extralarge": ModelSizeExtraLarge,
		"large":      ModelSizeLarge,
		"medium":     ModelSizeMedium,
		"none":       ModelSizeNone,
		"small":      ModelSizeSmall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ModelSize(input)
	return &out, nil
}

type MonitoringAlertNotificationType string

const (
	MonitoringAlertNotificationTypeAzureMonitor MonitoringAlertNotificationType = "AzureMonitor"
	MonitoringAlertNotificationTypeEmail        MonitoringAlertNotificationType = "Email"
)

func PossibleValuesForMonitoringAlertNotificationType() []string {
	return []string{
		string(MonitoringAlertNotificationTypeAzureMonitor),
		string(MonitoringAlertNotificationTypeEmail),
	}
}

func (s *MonitoringAlertNotificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitoringAlertNotificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitoringAlertNotificationType(input string) (*MonitoringAlertNotificationType, error) {
	vals := map[string]MonitoringAlertNotificationType{
		"azuremonitor": MonitoringAlertNotificationTypeAzureMonitor,
		"email":        MonitoringAlertNotificationTypeEmail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitoringAlertNotificationType(input)
	return &out, nil
}

type MonitoringFeatureDataType string

const (
	MonitoringFeatureDataTypeCategorical MonitoringFeatureDataType = "Categorical"
	MonitoringFeatureDataTypeNumerical   MonitoringFeatureDataType = "Numerical"
)

func PossibleValuesForMonitoringFeatureDataType() []string {
	return []string{
		string(MonitoringFeatureDataTypeCategorical),
		string(MonitoringFeatureDataTypeNumerical),
	}
}

func (s *MonitoringFeatureDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitoringFeatureDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitoringFeatureDataType(input string) (*MonitoringFeatureDataType, error) {
	vals := map[string]MonitoringFeatureDataType{
		"categorical": MonitoringFeatureDataTypeCategorical,
		"numerical":   MonitoringFeatureDataTypeNumerical,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitoringFeatureDataType(input)
	return &out, nil
}

type MonitoringFeatureFilterType string

const (
	MonitoringFeatureFilterTypeAllFeatures       MonitoringFeatureFilterType = "AllFeatures"
	MonitoringFeatureFilterTypeFeatureSubset     MonitoringFeatureFilterType = "FeatureSubset"
	MonitoringFeatureFilterTypeTopNByAttribution MonitoringFeatureFilterType = "TopNByAttribution"
)

func PossibleValuesForMonitoringFeatureFilterType() []string {
	return []string{
		string(MonitoringFeatureFilterTypeAllFeatures),
		string(MonitoringFeatureFilterTypeFeatureSubset),
		string(MonitoringFeatureFilterTypeTopNByAttribution),
	}
}

func (s *MonitoringFeatureFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitoringFeatureFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitoringFeatureFilterType(input string) (*MonitoringFeatureFilterType, error) {
	vals := map[string]MonitoringFeatureFilterType{
		"allfeatures":       MonitoringFeatureFilterTypeAllFeatures,
		"featuresubset":     MonitoringFeatureFilterTypeFeatureSubset,
		"topnbyattribution": MonitoringFeatureFilterTypeTopNByAttribution,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitoringFeatureFilterType(input)
	return &out, nil
}

type MonitoringInputDataContext string

const (
	MonitoringInputDataContextGroundTruth  MonitoringInputDataContext = "GroundTruth"
	MonitoringInputDataContextModelInputs  MonitoringInputDataContext = "ModelInputs"
	MonitoringInputDataContextModelOutputs MonitoringInputDataContext = "ModelOutputs"
	MonitoringInputDataContextTest         MonitoringInputDataContext = "Test"
	MonitoringInputDataContextTraining     MonitoringInputDataContext = "Training"
	MonitoringInputDataContextValidation   MonitoringInputDataContext = "Validation"
)

func PossibleValuesForMonitoringInputDataContext() []string {
	return []string{
		string(MonitoringInputDataContextGroundTruth),
		string(MonitoringInputDataContextModelInputs),
		string(MonitoringInputDataContextModelOutputs),
		string(MonitoringInputDataContextTest),
		string(MonitoringInputDataContextTraining),
		string(MonitoringInputDataContextValidation),
	}
}

func (s *MonitoringInputDataContext) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitoringInputDataContext(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitoringInputDataContext(input string) (*MonitoringInputDataContext, error) {
	vals := map[string]MonitoringInputDataContext{
		"groundtruth":  MonitoringInputDataContextGroundTruth,
		"modelinputs":  MonitoringInputDataContextModelInputs,
		"modeloutputs": MonitoringInputDataContextModelOutputs,
		"test":         MonitoringInputDataContextTest,
		"training":     MonitoringInputDataContextTraining,
		"validation":   MonitoringInputDataContextValidation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitoringInputDataContext(input)
	return &out, nil
}

type MonitoringModelType string

const (
	MonitoringModelTypeClassification MonitoringModelType = "Classification"
	MonitoringModelTypeRegression     MonitoringModelType = "Regression"
)

func PossibleValuesForMonitoringModelType() []string {
	return []string{
		string(MonitoringModelTypeClassification),
		string(MonitoringModelTypeRegression),
	}
}

func (s *MonitoringModelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitoringModelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitoringModelType(input string) (*MonitoringModelType, error) {
	vals := map[string]MonitoringModelType{
		"classification": MonitoringModelTypeClassification,
		"regression":     MonitoringModelTypeRegression,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitoringModelType(input)
	return &out, nil
}

type MonitoringNotificationMode string

const (
	MonitoringNotificationModeDisabled MonitoringNotificationMode = "Disabled"
	MonitoringNotificationModeEnabled  MonitoringNotificationMode = "Enabled"
)

func PossibleValuesForMonitoringNotificationMode() []string {
	return []string{
		string(MonitoringNotificationModeDisabled),
		string(MonitoringNotificationModeEnabled),
	}
}

func (s *MonitoringNotificationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitoringNotificationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitoringNotificationMode(input string) (*MonitoringNotificationMode, error) {
	vals := map[string]MonitoringNotificationMode{
		"disabled": MonitoringNotificationModeDisabled,
		"enabled":  MonitoringNotificationModeEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitoringNotificationMode(input)
	return &out, nil
}

type MonitoringSignalType string

const (
	MonitoringSignalTypeCustom                  MonitoringSignalType = "Custom"
	MonitoringSignalTypeDataDrift               MonitoringSignalType = "DataDrift"
	MonitoringSignalTypeDataQuality             MonitoringSignalType = "DataQuality"
	MonitoringSignalTypeFeatureAttributionDrift MonitoringSignalType = "FeatureAttributionDrift"
	MonitoringSignalTypeModelPerformance        MonitoringSignalType = "ModelPerformance"
	MonitoringSignalTypePredictionDrift         MonitoringSignalType = "PredictionDrift"
)

func PossibleValuesForMonitoringSignalType() []string {
	return []string{
		string(MonitoringSignalTypeCustom),
		string(MonitoringSignalTypeDataDrift),
		string(MonitoringSignalTypeDataQuality),
		string(MonitoringSignalTypeFeatureAttributionDrift),
		string(MonitoringSignalTypeModelPerformance),
		string(MonitoringSignalTypePredictionDrift),
	}
}

func (s *MonitoringSignalType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMonitoringSignalType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMonitoringSignalType(input string) (*MonitoringSignalType, error) {
	vals := map[string]MonitoringSignalType{
		"custom":                  MonitoringSignalTypeCustom,
		"datadrift":               MonitoringSignalTypeDataDrift,
		"dataquality":             MonitoringSignalTypeDataQuality,
		"featureattributiondrift": MonitoringSignalTypeFeatureAttributionDrift,
		"modelperformance":        MonitoringSignalTypeModelPerformance,
		"predictiondrift":         MonitoringSignalTypePredictionDrift,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MonitoringSignalType(input)
	return &out, nil
}

type MultiSelect string

const (
	MultiSelectDisabled MultiSelect = "Disabled"
	MultiSelectEnabled  MultiSelect = "Enabled"
)

func PossibleValuesForMultiSelect() []string {
	return []string{
		string(MultiSelectDisabled),
		string(MultiSelectEnabled),
	}
}

func (s *MultiSelect) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiSelect(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiSelect(input string) (*MultiSelect, error) {
	vals := map[string]MultiSelect{
		"disabled": MultiSelectDisabled,
		"enabled":  MultiSelectEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiSelect(input)
	return &out, nil
}

type NCrossValidationsMode string

const (
	NCrossValidationsModeAuto   NCrossValidationsMode = "Auto"
	NCrossValidationsModeCustom NCrossValidationsMode = "Custom"
)

func PossibleValuesForNCrossValidationsMode() []string {
	return []string{
		string(NCrossValidationsModeAuto),
		string(NCrossValidationsModeCustom),
	}
}

func (s *NCrossValidationsMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNCrossValidationsMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNCrossValidationsMode(input string) (*NCrossValidationsMode, error) {
	vals := map[string]NCrossValidationsMode{
		"auto":   NCrossValidationsModeAuto,
		"custom": NCrossValidationsModeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NCrossValidationsMode(input)
	return &out, nil
}

type NlpLearningRateScheduler string

const (
	NlpLearningRateSchedulerConstant           NlpLearningRateScheduler = "Constant"
	NlpLearningRateSchedulerConstantWithWarmup NlpLearningRateScheduler = "ConstantWithWarmup"
	NlpLearningRateSchedulerCosine             NlpLearningRateScheduler = "Cosine"
	NlpLearningRateSchedulerCosineWithRestarts NlpLearningRateScheduler = "CosineWithRestarts"
	NlpLearningRateSchedulerLinear             NlpLearningRateScheduler = "Linear"
	NlpLearningRateSchedulerNone               NlpLearningRateScheduler = "None"
	NlpLearningRateSchedulerPolynomial         NlpLearningRateScheduler = "Polynomial"
)

func PossibleValuesForNlpLearningRateScheduler() []string {
	return []string{
		string(NlpLearningRateSchedulerConstant),
		string(NlpLearningRateSchedulerConstantWithWarmup),
		string(NlpLearningRateSchedulerCosine),
		string(NlpLearningRateSchedulerCosineWithRestarts),
		string(NlpLearningRateSchedulerLinear),
		string(NlpLearningRateSchedulerNone),
		string(NlpLearningRateSchedulerPolynomial),
	}
}

func (s *NlpLearningRateScheduler) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNlpLearningRateScheduler(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNlpLearningRateScheduler(input string) (*NlpLearningRateScheduler, error) {
	vals := map[string]NlpLearningRateScheduler{
		"constant":           NlpLearningRateSchedulerConstant,
		"constantwithwarmup": NlpLearningRateSchedulerConstantWithWarmup,
		"cosine":             NlpLearningRateSchedulerCosine,
		"cosinewithrestarts": NlpLearningRateSchedulerCosineWithRestarts,
		"linear":             NlpLearningRateSchedulerLinear,
		"none":               NlpLearningRateSchedulerNone,
		"polynomial":         NlpLearningRateSchedulerPolynomial,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NlpLearningRateScheduler(input)
	return &out, nil
}

type NodesValueType string

const (
	NodesValueTypeAll    NodesValueType = "All"
	NodesValueTypeCustom NodesValueType = "Custom"
)

func PossibleValuesForNodesValueType() []string {
	return []string{
		string(NodesValueTypeAll),
		string(NodesValueTypeCustom),
	}
}

func (s *NodesValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNodesValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNodesValueType(input string) (*NodesValueType, error) {
	vals := map[string]NodesValueType{
		"all":    NodesValueTypeAll,
		"custom": NodesValueTypeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NodesValueType(input)
	return &out, nil
}

type NumericalDataDriftMetric string

const (
	NumericalDataDriftMetricJensenShannonDistance          NumericalDataDriftMetric = "JensenShannonDistance"
	NumericalDataDriftMetricNormalizedWassersteinDistance  NumericalDataDriftMetric = "NormalizedWassersteinDistance"
	NumericalDataDriftMetricPopulationStabilityIndex       NumericalDataDriftMetric = "PopulationStabilityIndex"
	NumericalDataDriftMetricTwoSampleKolmogorovSmirnovTest NumericalDataDriftMetric = "TwoSampleKolmogorovSmirnovTest"
)

func PossibleValuesForNumericalDataDriftMetric() []string {
	return []string{
		string(NumericalDataDriftMetricJensenShannonDistance),
		string(NumericalDataDriftMetricNormalizedWassersteinDistance),
		string(NumericalDataDriftMetricPopulationStabilityIndex),
		string(NumericalDataDriftMetricTwoSampleKolmogorovSmirnovTest),
	}
}

func (s *NumericalDataDriftMetric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNumericalDataDriftMetric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNumericalDataDriftMetric(input string) (*NumericalDataDriftMetric, error) {
	vals := map[string]NumericalDataDriftMetric{
		"jensenshannondistance":          NumericalDataDriftMetricJensenShannonDistance,
		"normalizedwassersteindistance":  NumericalDataDriftMetricNormalizedWassersteinDistance,
		"populationstabilityindex":       NumericalDataDriftMetricPopulationStabilityIndex,
		"twosamplekolmogorovsmirnovtest": NumericalDataDriftMetricTwoSampleKolmogorovSmirnovTest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NumericalDataDriftMetric(input)
	return &out, nil
}

type NumericalDataQualityMetric string

const (
	NumericalDataQualityMetricDataTypeErrorRate NumericalDataQualityMetric = "DataTypeErrorRate"
	NumericalDataQualityMetricNullValueRate     NumericalDataQualityMetric = "NullValueRate"
	NumericalDataQualityMetricOutOfBoundsRate   NumericalDataQualityMetric = "OutOfBoundsRate"
)

func PossibleValuesForNumericalDataQualityMetric() []string {
	return []string{
		string(NumericalDataQualityMetricDataTypeErrorRate),
		string(NumericalDataQualityMetricNullValueRate),
		string(NumericalDataQualityMetricOutOfBoundsRate),
	}
}

func (s *NumericalDataQualityMetric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNumericalDataQualityMetric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNumericalDataQualityMetric(input string) (*NumericalDataQualityMetric, error) {
	vals := map[string]NumericalDataQualityMetric{
		"datatypeerrorrate": NumericalDataQualityMetricDataTypeErrorRate,
		"nullvaluerate":     NumericalDataQualityMetricNullValueRate,
		"outofboundsrate":   NumericalDataQualityMetricOutOfBoundsRate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NumericalDataQualityMetric(input)
	return &out, nil
}

type NumericalPredictionDriftMetric string

const (
	NumericalPredictionDriftMetricJensenShannonDistance          NumericalPredictionDriftMetric = "JensenShannonDistance"
	NumericalPredictionDriftMetricNormalizedWassersteinDistance  NumericalPredictionDriftMetric = "NormalizedWassersteinDistance"
	NumericalPredictionDriftMetricPopulationStabilityIndex       NumericalPredictionDriftMetric = "PopulationStabilityIndex"
	NumericalPredictionDriftMetricTwoSampleKolmogorovSmirnovTest NumericalPredictionDriftMetric = "TwoSampleKolmogorovSmirnovTest"
)

func PossibleValuesForNumericalPredictionDriftMetric() []string {
	return []string{
		string(NumericalPredictionDriftMetricJensenShannonDistance),
		string(NumericalPredictionDriftMetricNormalizedWassersteinDistance),
		string(NumericalPredictionDriftMetricPopulationStabilityIndex),
		string(NumericalPredictionDriftMetricTwoSampleKolmogorovSmirnovTest),
	}
}

func (s *NumericalPredictionDriftMetric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNumericalPredictionDriftMetric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNumericalPredictionDriftMetric(input string) (*NumericalPredictionDriftMetric, error) {
	vals := map[string]NumericalPredictionDriftMetric{
		"jensenshannondistance":          NumericalPredictionDriftMetricJensenShannonDistance,
		"normalizedwassersteindistance":  NumericalPredictionDriftMetricNormalizedWassersteinDistance,
		"populationstabilityindex":       NumericalPredictionDriftMetricPopulationStabilityIndex,
		"twosamplekolmogorovsmirnovtest": NumericalPredictionDriftMetricTwoSampleKolmogorovSmirnovTest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NumericalPredictionDriftMetric(input)
	return &out, nil
}

type ObjectDetectionPrimaryMetrics string

const (
	ObjectDetectionPrimaryMetricsMeanAveragePrecision ObjectDetectionPrimaryMetrics = "MeanAveragePrecision"
)

func PossibleValuesForObjectDetectionPrimaryMetrics() []string {
	return []string{
		string(ObjectDetectionPrimaryMetricsMeanAveragePrecision),
	}
}

func (s *ObjectDetectionPrimaryMetrics) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectDetectionPrimaryMetrics(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectDetectionPrimaryMetrics(input string) (*ObjectDetectionPrimaryMetrics, error) {
	vals := map[string]ObjectDetectionPrimaryMetrics{
		"meanaverageprecision": ObjectDetectionPrimaryMetricsMeanAveragePrecision,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectDetectionPrimaryMetrics(input)
	return &out, nil
}

type OutputDeliveryMode string

const (
	OutputDeliveryModeDirect         OutputDeliveryMode = "Direct"
	OutputDeliveryModeReadWriteMount OutputDeliveryMode = "ReadWriteMount"
	OutputDeliveryModeUpload         OutputDeliveryMode = "Upload"
)

func PossibleValuesForOutputDeliveryMode() []string {
	return []string{
		string(OutputDeliveryModeDirect),
		string(OutputDeliveryModeReadWriteMount),
		string(OutputDeliveryModeUpload),
	}
}

func (s *OutputDeliveryMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutputDeliveryMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutputDeliveryMode(input string) (*OutputDeliveryMode, error) {
	vals := map[string]OutputDeliveryMode{
		"direct":         OutputDeliveryModeDirect,
		"readwritemount": OutputDeliveryModeReadWriteMount,
		"upload":         OutputDeliveryModeUpload,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputDeliveryMode(input)
	return &out, nil
}

type ProtectionLevel string

const (
	ProtectionLevelAll  ProtectionLevel = "All"
	ProtectionLevelNone ProtectionLevel = "None"
)

func PossibleValuesForProtectionLevel() []string {
	return []string{
		string(ProtectionLevelAll),
		string(ProtectionLevelNone),
	}
}

func (s *ProtectionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectionLevel(input string) (*ProtectionLevel, error) {
	vals := map[string]ProtectionLevel{
		"all":  ProtectionLevelAll,
		"none": ProtectionLevelNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionLevel(input)
	return &out, nil
}

type RandomSamplingAlgorithmRule string

const (
	RandomSamplingAlgorithmRuleRandom RandomSamplingAlgorithmRule = "Random"
	RandomSamplingAlgorithmRuleSobol  RandomSamplingAlgorithmRule = "Sobol"
)

func PossibleValuesForRandomSamplingAlgorithmRule() []string {
	return []string{
		string(RandomSamplingAlgorithmRuleRandom),
		string(RandomSamplingAlgorithmRuleSobol),
	}
}

func (s *RandomSamplingAlgorithmRule) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRandomSamplingAlgorithmRule(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRandomSamplingAlgorithmRule(input string) (*RandomSamplingAlgorithmRule, error) {
	vals := map[string]RandomSamplingAlgorithmRule{
		"random": RandomSamplingAlgorithmRuleRandom,
		"sobol":  RandomSamplingAlgorithmRuleSobol,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RandomSamplingAlgorithmRule(input)
	return &out, nil
}

type RecurrenceFrequency string

const (
	RecurrenceFrequencyDay    RecurrenceFrequency = "Day"
	RecurrenceFrequencyHour   RecurrenceFrequency = "Hour"
	RecurrenceFrequencyMinute RecurrenceFrequency = "Minute"
	RecurrenceFrequencyMonth  RecurrenceFrequency = "Month"
	RecurrenceFrequencyWeek   RecurrenceFrequency = "Week"
)

func PossibleValuesForRecurrenceFrequency() []string {
	return []string{
		string(RecurrenceFrequencyDay),
		string(RecurrenceFrequencyHour),
		string(RecurrenceFrequencyMinute),
		string(RecurrenceFrequencyMonth),
		string(RecurrenceFrequencyWeek),
	}
}

func (s *RecurrenceFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrenceFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrenceFrequency(input string) (*RecurrenceFrequency, error) {
	vals := map[string]RecurrenceFrequency{
		"day":    RecurrenceFrequencyDay,
		"hour":   RecurrenceFrequencyHour,
		"minute": RecurrenceFrequencyMinute,
		"month":  RecurrenceFrequencyMonth,
		"week":   RecurrenceFrequencyWeek,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrenceFrequency(input)
	return &out, nil
}

type RegressionModelPerformanceMetric string

const (
	RegressionModelPerformanceMetricMeanAbsoluteError    RegressionModelPerformanceMetric = "MeanAbsoluteError"
	RegressionModelPerformanceMetricMeanSquaredError     RegressionModelPerformanceMetric = "MeanSquaredError"
	RegressionModelPerformanceMetricRootMeanSquaredError RegressionModelPerformanceMetric = "RootMeanSquaredError"
)

func PossibleValuesForRegressionModelPerformanceMetric() []string {
	return []string{
		string(RegressionModelPerformanceMetricMeanAbsoluteError),
		string(RegressionModelPerformanceMetricMeanSquaredError),
		string(RegressionModelPerformanceMetricRootMeanSquaredError),
	}
}

func (s *RegressionModelPerformanceMetric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegressionModelPerformanceMetric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegressionModelPerformanceMetric(input string) (*RegressionModelPerformanceMetric, error) {
	vals := map[string]RegressionModelPerformanceMetric{
		"meanabsoluteerror":    RegressionModelPerformanceMetricMeanAbsoluteError,
		"meansquarederror":     RegressionModelPerformanceMetricMeanSquaredError,
		"rootmeansquarederror": RegressionModelPerformanceMetricRootMeanSquaredError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegressionModelPerformanceMetric(input)
	return &out, nil
}

type RegressionModels string

const (
	RegressionModelsDecisionTree       RegressionModels = "DecisionTree"
	RegressionModelsElasticNet         RegressionModels = "ElasticNet"
	RegressionModelsExtremeRandomTrees RegressionModels = "ExtremeRandomTrees"
	RegressionModelsGradientBoosting   RegressionModels = "GradientBoosting"
	RegressionModelsKNN                RegressionModels = "KNN"
	RegressionModelsLassoLars          RegressionModels = "LassoLars"
	RegressionModelsLightGBM           RegressionModels = "LightGBM"
	RegressionModelsRandomForest       RegressionModels = "RandomForest"
	RegressionModelsSGD                RegressionModels = "SGD"
	RegressionModelsXGBoostRegressor   RegressionModels = "XGBoostRegressor"
)

func PossibleValuesForRegressionModels() []string {
	return []string{
		string(RegressionModelsDecisionTree),
		string(RegressionModelsElasticNet),
		string(RegressionModelsExtremeRandomTrees),
		string(RegressionModelsGradientBoosting),
		string(RegressionModelsKNN),
		string(RegressionModelsLassoLars),
		string(RegressionModelsLightGBM),
		string(RegressionModelsRandomForest),
		string(RegressionModelsSGD),
		string(RegressionModelsXGBoostRegressor),
	}
}

func (s *RegressionModels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegressionModels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegressionModels(input string) (*RegressionModels, error) {
	vals := map[string]RegressionModels{
		"decisiontree":       RegressionModelsDecisionTree,
		"elasticnet":         RegressionModelsElasticNet,
		"extremerandomtrees": RegressionModelsExtremeRandomTrees,
		"gradientboosting":   RegressionModelsGradientBoosting,
		"knn":                RegressionModelsKNN,
		"lassolars":          RegressionModelsLassoLars,
		"lightgbm":           RegressionModelsLightGBM,
		"randomforest":       RegressionModelsRandomForest,
		"sgd":                RegressionModelsSGD,
		"xgboostregressor":   RegressionModelsXGBoostRegressor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegressionModels(input)
	return &out, nil
}

type RegressionPrimaryMetrics string

const (
	RegressionPrimaryMetricsNormalizedMeanAbsoluteError    RegressionPrimaryMetrics = "NormalizedMeanAbsoluteError"
	RegressionPrimaryMetricsNormalizedRootMeanSquaredError RegressionPrimaryMetrics = "NormalizedRootMeanSquaredError"
	RegressionPrimaryMetricsRTwoScore                      RegressionPrimaryMetrics = "R2Score"
	RegressionPrimaryMetricsSpearmanCorrelation            RegressionPrimaryMetrics = "SpearmanCorrelation"
)

func PossibleValuesForRegressionPrimaryMetrics() []string {
	return []string{
		string(RegressionPrimaryMetricsNormalizedMeanAbsoluteError),
		string(RegressionPrimaryMetricsNormalizedRootMeanSquaredError),
		string(RegressionPrimaryMetricsRTwoScore),
		string(RegressionPrimaryMetricsSpearmanCorrelation),
	}
}

func (s *RegressionPrimaryMetrics) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegressionPrimaryMetrics(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegressionPrimaryMetrics(input string) (*RegressionPrimaryMetrics, error) {
	vals := map[string]RegressionPrimaryMetrics{
		"normalizedmeanabsoluteerror":    RegressionPrimaryMetricsNormalizedMeanAbsoluteError,
		"normalizedrootmeansquarederror": RegressionPrimaryMetricsNormalizedRootMeanSquaredError,
		"r2score":                        RegressionPrimaryMetricsRTwoScore,
		"spearmancorrelation":            RegressionPrimaryMetricsSpearmanCorrelation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegressionPrimaryMetrics(input)
	return &out, nil
}

type SamplingAlgorithmType string

const (
	SamplingAlgorithmTypeBayesian SamplingAlgorithmType = "Bayesian"
	SamplingAlgorithmTypeGrid     SamplingAlgorithmType = "Grid"
	SamplingAlgorithmTypeRandom   SamplingAlgorithmType = "Random"
)

func PossibleValuesForSamplingAlgorithmType() []string {
	return []string{
		string(SamplingAlgorithmTypeBayesian),
		string(SamplingAlgorithmTypeGrid),
		string(SamplingAlgorithmTypeRandom),
	}
}

func (s *SamplingAlgorithmType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSamplingAlgorithmType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSamplingAlgorithmType(input string) (*SamplingAlgorithmType, error) {
	vals := map[string]SamplingAlgorithmType{
		"bayesian": SamplingAlgorithmTypeBayesian,
		"grid":     SamplingAlgorithmTypeGrid,
		"random":   SamplingAlgorithmTypeRandom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SamplingAlgorithmType(input)
	return &out, nil
}

type ScheduleActionType string

const (
	ScheduleActionTypeCreateJob           ScheduleActionType = "CreateJob"
	ScheduleActionTypeCreateMonitor       ScheduleActionType = "CreateMonitor"
	ScheduleActionTypeImportData          ScheduleActionType = "ImportData"
	ScheduleActionTypeInvokeBatchEndpoint ScheduleActionType = "InvokeBatchEndpoint"
)

func PossibleValuesForScheduleActionType() []string {
	return []string{
		string(ScheduleActionTypeCreateJob),
		string(ScheduleActionTypeCreateMonitor),
		string(ScheduleActionTypeImportData),
		string(ScheduleActionTypeInvokeBatchEndpoint),
	}
}

func (s *ScheduleActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleActionType(input string) (*ScheduleActionType, error) {
	vals := map[string]ScheduleActionType{
		"createjob":           ScheduleActionTypeCreateJob,
		"createmonitor":       ScheduleActionTypeCreateMonitor,
		"importdata":          ScheduleActionTypeImportData,
		"invokebatchendpoint": ScheduleActionTypeInvokeBatchEndpoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleActionType(input)
	return &out, nil
}

type ScheduleListViewType string

const (
	ScheduleListViewTypeAll          ScheduleListViewType = "All"
	ScheduleListViewTypeDisabledOnly ScheduleListViewType = "DisabledOnly"
	ScheduleListViewTypeEnabledOnly  ScheduleListViewType = "EnabledOnly"
)

func PossibleValuesForScheduleListViewType() []string {
	return []string{
		string(ScheduleListViewTypeAll),
		string(ScheduleListViewTypeDisabledOnly),
		string(ScheduleListViewTypeEnabledOnly),
	}
}

func (s *ScheduleListViewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleListViewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleListViewType(input string) (*ScheduleListViewType, error) {
	vals := map[string]ScheduleListViewType{
		"all":          ScheduleListViewTypeAll,
		"disabledonly": ScheduleListViewTypeDisabledOnly,
		"enabledonly":  ScheduleListViewTypeEnabledOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleListViewType(input)
	return &out, nil
}

type ScheduleProvisioningStatus string

const (
	ScheduleProvisioningStatusCanceled  ScheduleProvisioningStatus = "Canceled"
	ScheduleProvisioningStatusCreating  ScheduleProvisioningStatus = "Creating"
	ScheduleProvisioningStatusDeleting  ScheduleProvisioningStatus = "Deleting"
	ScheduleProvisioningStatusFailed    ScheduleProvisioningStatus = "Failed"
	ScheduleProvisioningStatusSucceeded ScheduleProvisioningStatus = "Succeeded"
	ScheduleProvisioningStatusUpdating  ScheduleProvisioningStatus = "Updating"
)

func PossibleValuesForScheduleProvisioningStatus() []string {
	return []string{
		string(ScheduleProvisioningStatusCanceled),
		string(ScheduleProvisioningStatusCreating),
		string(ScheduleProvisioningStatusDeleting),
		string(ScheduleProvisioningStatusFailed),
		string(ScheduleProvisioningStatusSucceeded),
		string(ScheduleProvisioningStatusUpdating),
	}
}

func (s *ScheduleProvisioningStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleProvisioningStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleProvisioningStatus(input string) (*ScheduleProvisioningStatus, error) {
	vals := map[string]ScheduleProvisioningStatus{
		"canceled":  ScheduleProvisioningStatusCanceled,
		"creating":  ScheduleProvisioningStatusCreating,
		"deleting":  ScheduleProvisioningStatusDeleting,
		"failed":    ScheduleProvisioningStatusFailed,
		"succeeded": ScheduleProvisioningStatusSucceeded,
		"updating":  ScheduleProvisioningStatusUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleProvisioningStatus(input)
	return &out, nil
}

type SeasonalityMode string

const (
	SeasonalityModeAuto   SeasonalityMode = "Auto"
	SeasonalityModeCustom SeasonalityMode = "Custom"
)

func PossibleValuesForSeasonalityMode() []string {
	return []string{
		string(SeasonalityModeAuto),
		string(SeasonalityModeCustom),
	}
}

func (s *SeasonalityMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSeasonalityMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSeasonalityMode(input string) (*SeasonalityMode, error) {
	vals := map[string]SeasonalityMode{
		"auto":   SeasonalityModeAuto,
		"custom": SeasonalityModeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SeasonalityMode(input)
	return &out, nil
}

type ShortSeriesHandlingConfiguration string

const (
	ShortSeriesHandlingConfigurationAuto ShortSeriesHandlingConfiguration = "Auto"
	ShortSeriesHandlingConfigurationDrop ShortSeriesHandlingConfiguration = "Drop"
	ShortSeriesHandlingConfigurationNone ShortSeriesHandlingConfiguration = "None"
	ShortSeriesHandlingConfigurationPad  ShortSeriesHandlingConfiguration = "Pad"
)

func PossibleValuesForShortSeriesHandlingConfiguration() []string {
	return []string{
		string(ShortSeriesHandlingConfigurationAuto),
		string(ShortSeriesHandlingConfigurationDrop),
		string(ShortSeriesHandlingConfigurationNone),
		string(ShortSeriesHandlingConfigurationPad),
	}
}

func (s *ShortSeriesHandlingConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShortSeriesHandlingConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShortSeriesHandlingConfiguration(input string) (*ShortSeriesHandlingConfiguration, error) {
	vals := map[string]ShortSeriesHandlingConfiguration{
		"auto": ShortSeriesHandlingConfigurationAuto,
		"drop": ShortSeriesHandlingConfigurationDrop,
		"none": ShortSeriesHandlingConfigurationNone,
		"pad":  ShortSeriesHandlingConfigurationPad,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShortSeriesHandlingConfiguration(input)
	return &out, nil
}

type SparkJobEntryType string

const (
	SparkJobEntryTypeSparkJobPythonEntry SparkJobEntryType = "SparkJobPythonEntry"
	SparkJobEntryTypeSparkJobScalaEntry  SparkJobEntryType = "SparkJobScalaEntry"
)

func PossibleValuesForSparkJobEntryType() []string {
	return []string{
		string(SparkJobEntryTypeSparkJobPythonEntry),
		string(SparkJobEntryTypeSparkJobScalaEntry),
	}
}

func (s *SparkJobEntryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSparkJobEntryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSparkJobEntryType(input string) (*SparkJobEntryType, error) {
	vals := map[string]SparkJobEntryType{
		"sparkjobpythonentry": SparkJobEntryTypeSparkJobPythonEntry,
		"sparkjobscalaentry":  SparkJobEntryTypeSparkJobScalaEntry,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SparkJobEntryType(input)
	return &out, nil
}

type StackMetaLearnerType string

const (
	StackMetaLearnerTypeElasticNet           StackMetaLearnerType = "ElasticNet"
	StackMetaLearnerTypeElasticNetCV         StackMetaLearnerType = "ElasticNetCV"
	StackMetaLearnerTypeLightGBMClassifier   StackMetaLearnerType = "LightGBMClassifier"
	StackMetaLearnerTypeLightGBMRegressor    StackMetaLearnerType = "LightGBMRegressor"
	StackMetaLearnerTypeLinearRegression     StackMetaLearnerType = "LinearRegression"
	StackMetaLearnerTypeLogisticRegression   StackMetaLearnerType = "LogisticRegression"
	StackMetaLearnerTypeLogisticRegressionCV StackMetaLearnerType = "LogisticRegressionCV"
	StackMetaLearnerTypeNone                 StackMetaLearnerType = "None"
)

func PossibleValuesForStackMetaLearnerType() []string {
	return []string{
		string(StackMetaLearnerTypeElasticNet),
		string(StackMetaLearnerTypeElasticNetCV),
		string(StackMetaLearnerTypeLightGBMClassifier),
		string(StackMetaLearnerTypeLightGBMRegressor),
		string(StackMetaLearnerTypeLinearRegression),
		string(StackMetaLearnerTypeLogisticRegression),
		string(StackMetaLearnerTypeLogisticRegressionCV),
		string(StackMetaLearnerTypeNone),
	}
}

func (s *StackMetaLearnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStackMetaLearnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStackMetaLearnerType(input string) (*StackMetaLearnerType, error) {
	vals := map[string]StackMetaLearnerType{
		"elasticnet":           StackMetaLearnerTypeElasticNet,
		"elasticnetcv":         StackMetaLearnerTypeElasticNetCV,
		"lightgbmclassifier":   StackMetaLearnerTypeLightGBMClassifier,
		"lightgbmregressor":    StackMetaLearnerTypeLightGBMRegressor,
		"linearregression":     StackMetaLearnerTypeLinearRegression,
		"logisticregression":   StackMetaLearnerTypeLogisticRegression,
		"logisticregressioncv": StackMetaLearnerTypeLogisticRegressionCV,
		"none":                 StackMetaLearnerTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StackMetaLearnerType(input)
	return &out, nil
}

type StatusMessageLevel string

const (
	StatusMessageLevelError       StatusMessageLevel = "Error"
	StatusMessageLevelInformation StatusMessageLevel = "Information"
	StatusMessageLevelWarning     StatusMessageLevel = "Warning"
)

func PossibleValuesForStatusMessageLevel() []string {
	return []string{
		string(StatusMessageLevelError),
		string(StatusMessageLevelInformation),
		string(StatusMessageLevelWarning),
	}
}

func (s *StatusMessageLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusMessageLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusMessageLevel(input string) (*StatusMessageLevel, error) {
	vals := map[string]StatusMessageLevel{
		"error":       StatusMessageLevelError,
		"information": StatusMessageLevelInformation,
		"warning":     StatusMessageLevelWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusMessageLevel(input)
	return &out, nil
}

type StochasticOptimizer string

const (
	StochasticOptimizerAdam  StochasticOptimizer = "Adam"
	StochasticOptimizerAdamw StochasticOptimizer = "Adamw"
	StochasticOptimizerNone  StochasticOptimizer = "None"
	StochasticOptimizerSgd   StochasticOptimizer = "Sgd"
)

func PossibleValuesForStochasticOptimizer() []string {
	return []string{
		string(StochasticOptimizerAdam),
		string(StochasticOptimizerAdamw),
		string(StochasticOptimizerNone),
		string(StochasticOptimizerSgd),
	}
}

func (s *StochasticOptimizer) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStochasticOptimizer(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStochasticOptimizer(input string) (*StochasticOptimizer, error) {
	vals := map[string]StochasticOptimizer{
		"adam":  StochasticOptimizerAdam,
		"adamw": StochasticOptimizerAdamw,
		"none":  StochasticOptimizerNone,
		"sgd":   StochasticOptimizerSgd,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StochasticOptimizer(input)
	return &out, nil
}

type TargetAggregationFunction string

const (
	TargetAggregationFunctionMax  TargetAggregationFunction = "Max"
	TargetAggregationFunctionMean TargetAggregationFunction = "Mean"
	TargetAggregationFunctionMin  TargetAggregationFunction = "Min"
	TargetAggregationFunctionNone TargetAggregationFunction = "None"
	TargetAggregationFunctionSum  TargetAggregationFunction = "Sum"
)

func PossibleValuesForTargetAggregationFunction() []string {
	return []string{
		string(TargetAggregationFunctionMax),
		string(TargetAggregationFunctionMean),
		string(TargetAggregationFunctionMin),
		string(TargetAggregationFunctionNone),
		string(TargetAggregationFunctionSum),
	}
}

func (s *TargetAggregationFunction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetAggregationFunction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetAggregationFunction(input string) (*TargetAggregationFunction, error) {
	vals := map[string]TargetAggregationFunction{
		"max":  TargetAggregationFunctionMax,
		"mean": TargetAggregationFunctionMean,
		"min":  TargetAggregationFunctionMin,
		"none": TargetAggregationFunctionNone,
		"sum":  TargetAggregationFunctionSum,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetAggregationFunction(input)
	return &out, nil
}

type TargetLagsMode string

const (
	TargetLagsModeAuto   TargetLagsMode = "Auto"
	TargetLagsModeCustom TargetLagsMode = "Custom"
)

func PossibleValuesForTargetLagsMode() []string {
	return []string{
		string(TargetLagsModeAuto),
		string(TargetLagsModeCustom),
	}
}

func (s *TargetLagsMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetLagsMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetLagsMode(input string) (*TargetLagsMode, error) {
	vals := map[string]TargetLagsMode{
		"auto":   TargetLagsModeAuto,
		"custom": TargetLagsModeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetLagsMode(input)
	return &out, nil
}

type TargetRollingWindowSizeMode string

const (
	TargetRollingWindowSizeModeAuto   TargetRollingWindowSizeMode = "Auto"
	TargetRollingWindowSizeModeCustom TargetRollingWindowSizeMode = "Custom"
)

func PossibleValuesForTargetRollingWindowSizeMode() []string {
	return []string{
		string(TargetRollingWindowSizeModeAuto),
		string(TargetRollingWindowSizeModeCustom),
	}
}

func (s *TargetRollingWindowSizeMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetRollingWindowSizeMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetRollingWindowSizeMode(input string) (*TargetRollingWindowSizeMode, error) {
	vals := map[string]TargetRollingWindowSizeMode{
		"auto":   TargetRollingWindowSizeModeAuto,
		"custom": TargetRollingWindowSizeModeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetRollingWindowSizeMode(input)
	return &out, nil
}

type TaskType string

const (
	TaskTypeClassification                TaskType = "Classification"
	TaskTypeForecasting                   TaskType = "Forecasting"
	TaskTypeImageClassification           TaskType = "ImageClassification"
	TaskTypeImageClassificationMultilabel TaskType = "ImageClassificationMultilabel"
	TaskTypeImageInstanceSegmentation     TaskType = "ImageInstanceSegmentation"
	TaskTypeImageObjectDetection          TaskType = "ImageObjectDetection"
	TaskTypeRegression                    TaskType = "Regression"
	TaskTypeTextClassification            TaskType = "TextClassification"
	TaskTypeTextClassificationMultilabel  TaskType = "TextClassificationMultilabel"
	TaskTypeTextNER                       TaskType = "TextNER"
)

func PossibleValuesForTaskType() []string {
	return []string{
		string(TaskTypeClassification),
		string(TaskTypeForecasting),
		string(TaskTypeImageClassification),
		string(TaskTypeImageClassificationMultilabel),
		string(TaskTypeImageInstanceSegmentation),
		string(TaskTypeImageObjectDetection),
		string(TaskTypeRegression),
		string(TaskTypeTextClassification),
		string(TaskTypeTextClassificationMultilabel),
		string(TaskTypeTextNER),
	}
}

func (s *TaskType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaskType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaskType(input string) (*TaskType, error) {
	vals := map[string]TaskType{
		"classification":                TaskTypeClassification,
		"forecasting":                   TaskTypeForecasting,
		"imageclassification":           TaskTypeImageClassification,
		"imageclassificationmultilabel": TaskTypeImageClassificationMultilabel,
		"imageinstancesegmentation":     TaskTypeImageInstanceSegmentation,
		"imageobjectdetection":          TaskTypeImageObjectDetection,
		"regression":                    TaskTypeRegression,
		"textclassification":            TaskTypeTextClassification,
		"textclassificationmultilabel":  TaskTypeTextClassificationMultilabel,
		"textner":                       TaskTypeTextNER,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaskType(input)
	return &out, nil
}

type TextAnnotationType string

const (
	TextAnnotationTypeClassification         TextAnnotationType = "Classification"
	TextAnnotationTypeNamedEntityRecognition TextAnnotationType = "NamedEntityRecognition"
)

func PossibleValuesForTextAnnotationType() []string {
	return []string{
		string(TextAnnotationTypeClassification),
		string(TextAnnotationTypeNamedEntityRecognition),
	}
}

func (s *TextAnnotationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTextAnnotationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTextAnnotationType(input string) (*TextAnnotationType, error) {
	vals := map[string]TextAnnotationType{
		"classification":         TextAnnotationTypeClassification,
		"namedentityrecognition": TextAnnotationTypeNamedEntityRecognition,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TextAnnotationType(input)
	return &out, nil
}

type TrainingMode string

const (
	TrainingModeAuto           TrainingMode = "Auto"
	TrainingModeDistributed    TrainingMode = "Distributed"
	TrainingModeNonDistributed TrainingMode = "NonDistributed"
)

func PossibleValuesForTrainingMode() []string {
	return []string{
		string(TrainingModeAuto),
		string(TrainingModeDistributed),
		string(TrainingModeNonDistributed),
	}
}

func (s *TrainingMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingMode(input string) (*TrainingMode, error) {
	vals := map[string]TrainingMode{
		"auto":           TrainingModeAuto,
		"distributed":    TrainingModeDistributed,
		"nondistributed": TrainingModeNonDistributed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingMode(input)
	return &out, nil
}

type TriggerType string

const (
	TriggerTypeCron       TriggerType = "Cron"
	TriggerTypeRecurrence TriggerType = "Recurrence"
)

func PossibleValuesForTriggerType() []string {
	return []string{
		string(TriggerTypeCron),
		string(TriggerTypeRecurrence),
	}
}

func (s *TriggerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerType(input string) (*TriggerType, error) {
	vals := map[string]TriggerType{
		"cron":       TriggerTypeCron,
		"recurrence": TriggerTypeRecurrence,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerType(input)
	return &out, nil
}

type UseStl string

const (
	UseStlNone        UseStl = "None"
	UseStlSeason      UseStl = "Season"
	UseStlSeasonTrend UseStl = "SeasonTrend"
)

func PossibleValuesForUseStl() []string {
	return []string{
		string(UseStlNone),
		string(UseStlSeason),
		string(UseStlSeasonTrend),
	}
}

func (s *UseStl) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUseStl(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUseStl(input string) (*UseStl, error) {
	vals := map[string]UseStl{
		"none":        UseStlNone,
		"season":      UseStlSeason,
		"seasontrend": UseStlSeasonTrend,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UseStl(input)
	return &out, nil
}

type ValidationMetricType string

const (
	ValidationMetricTypeCoco    ValidationMetricType = "Coco"
	ValidationMetricTypeCocoVoc ValidationMetricType = "CocoVoc"
	ValidationMetricTypeNone    ValidationMetricType = "None"
	ValidationMetricTypeVoc     ValidationMetricType = "Voc"
)

func PossibleValuesForValidationMetricType() []string {
	return []string{
		string(ValidationMetricTypeCoco),
		string(ValidationMetricTypeCocoVoc),
		string(ValidationMetricTypeNone),
		string(ValidationMetricTypeVoc),
	}
}

func (s *ValidationMetricType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseValidationMetricType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseValidationMetricType(input string) (*ValidationMetricType, error) {
	vals := map[string]ValidationMetricType{
		"coco":    ValidationMetricTypeCoco,
		"cocovoc": ValidationMetricTypeCocoVoc,
		"none":    ValidationMetricTypeNone,
		"voc":     ValidationMetricTypeVoc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValidationMetricType(input)
	return &out, nil
}

type WebhookType string

const (
	WebhookTypeAzureDevOps WebhookType = "AzureDevOps"
)

func PossibleValuesForWebhookType() []string {
	return []string{
		string(WebhookTypeAzureDevOps),
	}
}

func (s *WebhookType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebhookType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebhookType(input string) (*WebhookType, error) {
	vals := map[string]WebhookType{
		"azuredevops": WebhookTypeAzureDevOps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebhookType(input)
	return &out, nil
}

type WeekDay string

const (
	WeekDayFriday    WeekDay = "Friday"
	WeekDayMonday    WeekDay = "Monday"
	WeekDaySaturday  WeekDay = "Saturday"
	WeekDaySunday    WeekDay = "Sunday"
	WeekDayThursday  WeekDay = "Thursday"
	WeekDayTuesday   WeekDay = "Tuesday"
	WeekDayWednesday WeekDay = "Wednesday"
)

func PossibleValuesForWeekDay() []string {
	return []string{
		string(WeekDayFriday),
		string(WeekDayMonday),
		string(WeekDaySaturday),
		string(WeekDaySunday),
		string(WeekDayThursday),
		string(WeekDayTuesday),
		string(WeekDayWednesday),
	}
}

func (s *WeekDay) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWeekDay(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWeekDay(input string) (*WeekDay, error) {
	vals := map[string]WeekDay{
		"friday":    WeekDayFriday,
		"monday":    WeekDayMonday,
		"saturday":  WeekDaySaturday,
		"sunday":    WeekDaySunday,
		"thursday":  WeekDayThursday,
		"tuesday":   WeekDayTuesday,
		"wednesday": WeekDayWednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WeekDay(input)
	return &out, nil
}
