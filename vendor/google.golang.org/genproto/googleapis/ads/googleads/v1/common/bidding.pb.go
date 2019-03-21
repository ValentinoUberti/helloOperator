// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/bidding.proto

package common // import "google.golang.org/genproto/googleapis/ads/googleads/v1/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An automated bidding strategy that raises bids for clicks
// that seem more likely to lead to a conversion and lowers
// them for clicks where they seem less likely.
type EnhancedCpc struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnhancedCpc) Reset()         { *m = EnhancedCpc{} }
func (m *EnhancedCpc) String() string { return proto.CompactTextString(m) }
func (*EnhancedCpc) ProtoMessage()    {}
func (*EnhancedCpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{0}
}
func (m *EnhancedCpc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnhancedCpc.Unmarshal(m, b)
}
func (m *EnhancedCpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnhancedCpc.Marshal(b, m, deterministic)
}
func (dst *EnhancedCpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnhancedCpc.Merge(dst, src)
}
func (m *EnhancedCpc) XXX_Size() int {
	return xxx_messageInfo_EnhancedCpc.Size(m)
}
func (m *EnhancedCpc) XXX_DiscardUnknown() {
	xxx_messageInfo_EnhancedCpc.DiscardUnknown(m)
}

var xxx_messageInfo_EnhancedCpc proto.InternalMessageInfo

// Manual click-based bidding where user pays per click.
type ManualCpc struct {
	// Whether bids are to be enhanced based on conversion optimizer data.
	EnhancedCpcEnabled   *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enhanced_cpc_enabled,json=enhancedCpcEnabled,proto3" json:"enhanced_cpc_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ManualCpc) Reset()         { *m = ManualCpc{} }
func (m *ManualCpc) String() string { return proto.CompactTextString(m) }
func (*ManualCpc) ProtoMessage()    {}
func (*ManualCpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{1}
}
func (m *ManualCpc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualCpc.Unmarshal(m, b)
}
func (m *ManualCpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualCpc.Marshal(b, m, deterministic)
}
func (dst *ManualCpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualCpc.Merge(dst, src)
}
func (m *ManualCpc) XXX_Size() int {
	return xxx_messageInfo_ManualCpc.Size(m)
}
func (m *ManualCpc) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualCpc.DiscardUnknown(m)
}

var xxx_messageInfo_ManualCpc proto.InternalMessageInfo

func (m *ManualCpc) GetEnhancedCpcEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.EnhancedCpcEnabled
	}
	return nil
}

// Manual impression-based bidding where user pays per thousand impressions.
type ManualCpm struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManualCpm) Reset()         { *m = ManualCpm{} }
func (m *ManualCpm) String() string { return proto.CompactTextString(m) }
func (*ManualCpm) ProtoMessage()    {}
func (*ManualCpm) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{2}
}
func (m *ManualCpm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualCpm.Unmarshal(m, b)
}
func (m *ManualCpm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualCpm.Marshal(b, m, deterministic)
}
func (dst *ManualCpm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualCpm.Merge(dst, src)
}
func (m *ManualCpm) XXX_Size() int {
	return xxx_messageInfo_ManualCpm.Size(m)
}
func (m *ManualCpm) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualCpm.DiscardUnknown(m)
}

var xxx_messageInfo_ManualCpm proto.InternalMessageInfo

// View based bidding where user pays per video view.
type ManualCpv struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManualCpv) Reset()         { *m = ManualCpv{} }
func (m *ManualCpv) String() string { return proto.CompactTextString(m) }
func (*ManualCpv) ProtoMessage()    {}
func (*ManualCpv) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{3}
}
func (m *ManualCpv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualCpv.Unmarshal(m, b)
}
func (m *ManualCpv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualCpv.Marshal(b, m, deterministic)
}
func (dst *ManualCpv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualCpv.Merge(dst, src)
}
func (m *ManualCpv) XXX_Size() int {
	return xxx_messageInfo_ManualCpv.Size(m)
}
func (m *ManualCpv) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualCpv.DiscardUnknown(m)
}

var xxx_messageInfo_ManualCpv proto.InternalMessageInfo

// An automated bidding strategy that sets bids to help get the most conversions
// for your campaign while spending your budget.
type MaximizeConversions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaximizeConversions) Reset()         { *m = MaximizeConversions{} }
func (m *MaximizeConversions) String() string { return proto.CompactTextString(m) }
func (*MaximizeConversions) ProtoMessage()    {}
func (*MaximizeConversions) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{4}
}
func (m *MaximizeConversions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaximizeConversions.Unmarshal(m, b)
}
func (m *MaximizeConversions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaximizeConversions.Marshal(b, m, deterministic)
}
func (dst *MaximizeConversions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaximizeConversions.Merge(dst, src)
}
func (m *MaximizeConversions) XXX_Size() int {
	return xxx_messageInfo_MaximizeConversions.Size(m)
}
func (m *MaximizeConversions) XXX_DiscardUnknown() {
	xxx_messageInfo_MaximizeConversions.DiscardUnknown(m)
}

var xxx_messageInfo_MaximizeConversions proto.InternalMessageInfo

// An automated bidding strategy which tries to maximize conversion value
// given a daily budget.
type MaximizeConversionValue struct {
	// The target return on ad spend (ROAS) option. If set, the bid strategy will
	// maximize revenue while averaging the target return on ad spend. If the
	// target ROAS is high, the bid strategy may not be able to spend the full
	// budget. If the target ROAS is not set, the bid strategy will aim to
	// achieve the highest possible ROAS for the budget.
	TargetRoas           *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MaximizeConversionValue) Reset()         { *m = MaximizeConversionValue{} }
func (m *MaximizeConversionValue) String() string { return proto.CompactTextString(m) }
func (*MaximizeConversionValue) ProtoMessage()    {}
func (*MaximizeConversionValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{5}
}
func (m *MaximizeConversionValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaximizeConversionValue.Unmarshal(m, b)
}
func (m *MaximizeConversionValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaximizeConversionValue.Marshal(b, m, deterministic)
}
func (dst *MaximizeConversionValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaximizeConversionValue.Merge(dst, src)
}
func (m *MaximizeConversionValue) XXX_Size() int {
	return xxx_messageInfo_MaximizeConversionValue.Size(m)
}
func (m *MaximizeConversionValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MaximizeConversionValue.DiscardUnknown(m)
}

var xxx_messageInfo_MaximizeConversionValue proto.InternalMessageInfo

func (m *MaximizeConversionValue) GetTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.TargetRoas
	}
	return nil
}

// An automated bidding strategy which sets CPC bids to target impressions on
// page one, or page one promoted slots on google.com.
type PageOnePromoted struct {
	// The strategy goal of where impressions are desired to be shown on
	// search result pages.
	StrategyGoal enums.PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal `protobuf:"varint,1,opt,name=strategy_goal,json=strategyGoal,proto3,enum=google.ads.googleads.v1.enums.PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal" json:"strategy_goal,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Bid multiplier to be applied to the relevant bid estimate (depending on
	// the `strategy_goal`) in determining a keyword's new CPC bid.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// Whether the strategy should always follow bid estimate changes, or only
	// increase.
	// If false, always sets a keyword's new bid to the current bid estimate.
	// If true, only updates a keyword's bid if the current bid estimate is
	// greater than the current bid.
	OnlyRaiseCpcBids *wrappers.BoolValue `protobuf:"bytes,4,opt,name=only_raise_cpc_bids,json=onlyRaiseCpcBids,proto3" json:"only_raise_cpc_bids,omitempty"`
	// Whether the strategy is allowed to raise bids when the throttling
	// rate of the budget it is serving out of rises above a threshold.
	RaiseCpcBidWhenBudgetConstrained *wrappers.BoolValue `protobuf:"bytes,5,opt,name=raise_cpc_bid_when_budget_constrained,json=raiseCpcBidWhenBudgetConstrained,proto3" json:"raise_cpc_bid_when_budget_constrained,omitempty"`
	// Whether the strategy is allowed to raise bids on keywords with
	// lower-range quality scores.
	RaiseCpcBidWhenQualityScoreIsLow *wrappers.BoolValue `protobuf:"bytes,6,opt,name=raise_cpc_bid_when_quality_score_is_low,json=raiseCpcBidWhenQualityScoreIsLow,proto3" json:"raise_cpc_bid_when_quality_score_is_low,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}            `json:"-"`
	XXX_unrecognized                 []byte              `json:"-"`
	XXX_sizecache                    int32               `json:"-"`
}

func (m *PageOnePromoted) Reset()         { *m = PageOnePromoted{} }
func (m *PageOnePromoted) String() string { return proto.CompactTextString(m) }
func (*PageOnePromoted) ProtoMessage()    {}
func (*PageOnePromoted) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{6}
}
func (m *PageOnePromoted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageOnePromoted.Unmarshal(m, b)
}
func (m *PageOnePromoted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageOnePromoted.Marshal(b, m, deterministic)
}
func (dst *PageOnePromoted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageOnePromoted.Merge(dst, src)
}
func (m *PageOnePromoted) XXX_Size() int {
	return xxx_messageInfo_PageOnePromoted.Size(m)
}
func (m *PageOnePromoted) XXX_DiscardUnknown() {
	xxx_messageInfo_PageOnePromoted.DiscardUnknown(m)
}

var xxx_messageInfo_PageOnePromoted proto.InternalMessageInfo

func (m *PageOnePromoted) GetStrategyGoal() enums.PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal {
	if m != nil {
		return m.StrategyGoal
	}
	return enums.PageOnePromotedStrategyGoalEnum_UNSPECIFIED
}

func (m *PageOnePromoted) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *PageOnePromoted) GetBidModifier() *wrappers.DoubleValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *PageOnePromoted) GetOnlyRaiseCpcBids() *wrappers.BoolValue {
	if m != nil {
		return m.OnlyRaiseCpcBids
	}
	return nil
}

func (m *PageOnePromoted) GetRaiseCpcBidWhenBudgetConstrained() *wrappers.BoolValue {
	if m != nil {
		return m.RaiseCpcBidWhenBudgetConstrained
	}
	return nil
}

func (m *PageOnePromoted) GetRaiseCpcBidWhenQualityScoreIsLow() *wrappers.BoolValue {
	if m != nil {
		return m.RaiseCpcBidWhenQualityScoreIsLow
	}
	return nil
}

// An automated bid strategy that sets bids to help get as many conversions as
// possible at the target cost-per-acquisition (CPA) you set.
type TargetCpa struct {
	// Average CPA target.
	// This target should be greater than or equal to minimum billable unit based
	// on the currency for the account.
	TargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=target_cpa_micros,json=targetCpaMicros,proto3" json:"target_cpa_micros,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Minimum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidFloorMicros    *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_floor_micros,json=cpcBidFloorMicros,proto3" json:"cpc_bid_floor_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetCpa) Reset()         { *m = TargetCpa{} }
func (m *TargetCpa) String() string { return proto.CompactTextString(m) }
func (*TargetCpa) ProtoMessage()    {}
func (*TargetCpa) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{7}
}
func (m *TargetCpa) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCpa.Unmarshal(m, b)
}
func (m *TargetCpa) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCpa.Marshal(b, m, deterministic)
}
func (dst *TargetCpa) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCpa.Merge(dst, src)
}
func (m *TargetCpa) XXX_Size() int {
	return xxx_messageInfo_TargetCpa.Size(m)
}
func (m *TargetCpa) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCpa.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCpa proto.InternalMessageInfo

func (m *TargetCpa) GetTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpaMicros
	}
	return nil
}

func (m *TargetCpa) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *TargetCpa) GetCpcBidFloorMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidFloorMicros
	}
	return nil
}

// Target CPM (cost per thousand impressions) is an automated bidding strategy
// that sets bids to optimize performance given the target CPM you set.
type TargetCpm struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetCpm) Reset()         { *m = TargetCpm{} }
func (m *TargetCpm) String() string { return proto.CompactTextString(m) }
func (*TargetCpm) ProtoMessage()    {}
func (*TargetCpm) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{8}
}
func (m *TargetCpm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCpm.Unmarshal(m, b)
}
func (m *TargetCpm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCpm.Marshal(b, m, deterministic)
}
func (dst *TargetCpm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCpm.Merge(dst, src)
}
func (m *TargetCpm) XXX_Size() int {
	return xxx_messageInfo_TargetCpm.Size(m)
}
func (m *TargetCpm) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCpm.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCpm proto.InternalMessageInfo

// An automated bidding strategy that sets bids so that a certain percentage of
// search ads are shown at the top of the first page (or other targeted
// location).
// Next Id = 4
type TargetImpressionShare struct {
	// The targeted location on the search results page.
	Location enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation `protobuf:"varint,1,opt,name=location,proto3,enum=google.ads.googleads.v1.enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation" json:"location,omitempty"`
	// The desired fraction of ads to be shown in the targeted location in micros.
	// E.g. 1% equals 10,000.
	LocationFractionMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=location_fraction_micros,json=locationFractionMicros,proto3" json:"location_fraction_micros,omitempty"`
	// The highest CPC bid the automated bidding system is permitted to specify.
	// This is a required field entered by the advertiser that sets the ceiling
	// and specified in local micros.
	CpcBidCeilingMicros  *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetImpressionShare) Reset()         { *m = TargetImpressionShare{} }
func (m *TargetImpressionShare) String() string { return proto.CompactTextString(m) }
func (*TargetImpressionShare) ProtoMessage()    {}
func (*TargetImpressionShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{9}
}
func (m *TargetImpressionShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetImpressionShare.Unmarshal(m, b)
}
func (m *TargetImpressionShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetImpressionShare.Marshal(b, m, deterministic)
}
func (dst *TargetImpressionShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetImpressionShare.Merge(dst, src)
}
func (m *TargetImpressionShare) XXX_Size() int {
	return xxx_messageInfo_TargetImpressionShare.Size(m)
}
func (m *TargetImpressionShare) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetImpressionShare.DiscardUnknown(m)
}

var xxx_messageInfo_TargetImpressionShare proto.InternalMessageInfo

func (m *TargetImpressionShare) GetLocation() enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation {
	if m != nil {
		return m.Location
	}
	return enums.TargetImpressionShareLocationEnum_UNSPECIFIED
}

func (m *TargetImpressionShare) GetLocationFractionMicros() *wrappers.Int64Value {
	if m != nil {
		return m.LocationFractionMicros
	}
	return nil
}

func (m *TargetImpressionShare) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

// An automated bidding strategy that sets bids based on the target fraction of
// auctions where the advertiser should outrank a specific competitor.
type TargetOutrankShare struct {
	// The target fraction of auctions where the advertiser should outrank the
	// competitor.
	// The advertiser outranks the competitor in an auction if either the
	// advertiser appears above the competitor in the search results, or appears
	// in the search results when the competitor does not.
	// Value must be between 1 and 1000000, inclusive.
	TargetOutrankShareMicros *wrappers.Int32Value `protobuf:"bytes,1,opt,name=target_outrank_share_micros,json=targetOutrankShareMicros,proto3" json:"target_outrank_share_micros,omitempty"`
	// Competitor's visible domain URL.
	CompetitorDomain *wrappers.StringValue `protobuf:"bytes,2,opt,name=competitor_domain,json=competitorDomain,proto3" json:"competitor_domain,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Whether the strategy should always follow bid estimate changes,
	// or only increase.
	// If false, always set a keyword's new bid to the current bid estimate.
	// If true, only updates a keyword's bid if the current bid estimate is
	// greater than the current bid.
	OnlyRaiseCpcBids *wrappers.BoolValue `protobuf:"bytes,4,opt,name=only_raise_cpc_bids,json=onlyRaiseCpcBids,proto3" json:"only_raise_cpc_bids,omitempty"`
	// Whether the strategy is allowed to raise bids on keywords with
	// lower-range quality scores.
	RaiseCpcBidWhenQualityScoreIsLow *wrappers.BoolValue `protobuf:"bytes,5,opt,name=raise_cpc_bid_when_quality_score_is_low,json=raiseCpcBidWhenQualityScoreIsLow,proto3" json:"raise_cpc_bid_when_quality_score_is_low,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}            `json:"-"`
	XXX_unrecognized                 []byte              `json:"-"`
	XXX_sizecache                    int32               `json:"-"`
}

func (m *TargetOutrankShare) Reset()         { *m = TargetOutrankShare{} }
func (m *TargetOutrankShare) String() string { return proto.CompactTextString(m) }
func (*TargetOutrankShare) ProtoMessage()    {}
func (*TargetOutrankShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{10}
}
func (m *TargetOutrankShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetOutrankShare.Unmarshal(m, b)
}
func (m *TargetOutrankShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetOutrankShare.Marshal(b, m, deterministic)
}
func (dst *TargetOutrankShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetOutrankShare.Merge(dst, src)
}
func (m *TargetOutrankShare) XXX_Size() int {
	return xxx_messageInfo_TargetOutrankShare.Size(m)
}
func (m *TargetOutrankShare) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetOutrankShare.DiscardUnknown(m)
}

var xxx_messageInfo_TargetOutrankShare proto.InternalMessageInfo

func (m *TargetOutrankShare) GetTargetOutrankShareMicros() *wrappers.Int32Value {
	if m != nil {
		return m.TargetOutrankShareMicros
	}
	return nil
}

func (m *TargetOutrankShare) GetCompetitorDomain() *wrappers.StringValue {
	if m != nil {
		return m.CompetitorDomain
	}
	return nil
}

func (m *TargetOutrankShare) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *TargetOutrankShare) GetOnlyRaiseCpcBids() *wrappers.BoolValue {
	if m != nil {
		return m.OnlyRaiseCpcBids
	}
	return nil
}

func (m *TargetOutrankShare) GetRaiseCpcBidWhenQualityScoreIsLow() *wrappers.BoolValue {
	if m != nil {
		return m.RaiseCpcBidWhenQualityScoreIsLow
	}
	return nil
}

// An automated bidding strategy that helps you maximize revenue while
// averaging a specific target return on ad spend (ROAS).
type TargetRoas struct {
	// Required. The desired revenue (based on conversion data) per unit of spend.
	// Value must be between 0.01 and 1000.0, inclusive.
	TargetRoas *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Minimum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidFloorMicros    *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_floor_micros,json=cpcBidFloorMicros,proto3" json:"cpc_bid_floor_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetRoas) Reset()         { *m = TargetRoas{} }
func (m *TargetRoas) String() string { return proto.CompactTextString(m) }
func (*TargetRoas) ProtoMessage()    {}
func (*TargetRoas) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{11}
}
func (m *TargetRoas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetRoas.Unmarshal(m, b)
}
func (m *TargetRoas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetRoas.Marshal(b, m, deterministic)
}
func (dst *TargetRoas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetRoas.Merge(dst, src)
}
func (m *TargetRoas) XXX_Size() int {
	return xxx_messageInfo_TargetRoas.Size(m)
}
func (m *TargetRoas) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetRoas.DiscardUnknown(m)
}

var xxx_messageInfo_TargetRoas proto.InternalMessageInfo

func (m *TargetRoas) GetTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.TargetRoas
	}
	return nil
}

func (m *TargetRoas) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *TargetRoas) GetCpcBidFloorMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidFloorMicros
	}
	return nil
}

// An automated bid strategy that sets your bids to help get as many clicks
// as possible within your budget.
type TargetSpend struct {
	// The spend target under which to maximize clicks.
	// A TargetSpend bidder will attempt to spend the smaller of this value
	// or the natural throttling spend amount.
	// If not specified, the budget is used as the spend target.
	TargetSpendMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=target_spend_micros,json=targetSpendMicros,proto3" json:"target_spend_micros,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros  *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetSpend) Reset()         { *m = TargetSpend{} }
func (m *TargetSpend) String() string { return proto.CompactTextString(m) }
func (*TargetSpend) ProtoMessage()    {}
func (*TargetSpend) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{12}
}
func (m *TargetSpend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetSpend.Unmarshal(m, b)
}
func (m *TargetSpend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetSpend.Marshal(b, m, deterministic)
}
func (dst *TargetSpend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetSpend.Merge(dst, src)
}
func (m *TargetSpend) XXX_Size() int {
	return xxx_messageInfo_TargetSpend.Size(m)
}
func (m *TargetSpend) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetSpend.DiscardUnknown(m)
}

var xxx_messageInfo_TargetSpend proto.InternalMessageInfo

func (m *TargetSpend) GetTargetSpendMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetSpendMicros
	}
	return nil
}

func (m *TargetSpend) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

// A bidding strategy where bids are a fraction of the advertised price for
// some good or service.
type PercentCpc struct {
	// Maximum bid limit that can be set by the bid strategy. This is
	// an optional field entered by the advertiser and specified in local micros.
	// Note: A zero value is interpreted in the same way as having bid_ceiling
	// undefined.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Adjusts the bid for each auction upward or downward, depending on the
	// likelihood of a conversion. Individual bids may exceed
	// cpc_bid_ceiling_micros, but the average bid amount for a campaign should
	// not.
	EnhancedCpcEnabled   *wrappers.BoolValue `protobuf:"bytes,2,opt,name=enhanced_cpc_enabled,json=enhancedCpcEnabled,proto3" json:"enhanced_cpc_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PercentCpc) Reset()         { *m = PercentCpc{} }
func (m *PercentCpc) String() string { return proto.CompactTextString(m) }
func (*PercentCpc) ProtoMessage()    {}
func (*PercentCpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_bidding_ebc1284101877700, []int{13}
}
func (m *PercentCpc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PercentCpc.Unmarshal(m, b)
}
func (m *PercentCpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PercentCpc.Marshal(b, m, deterministic)
}
func (dst *PercentCpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PercentCpc.Merge(dst, src)
}
func (m *PercentCpc) XXX_Size() int {
	return xxx_messageInfo_PercentCpc.Size(m)
}
func (m *PercentCpc) XXX_DiscardUnknown() {
	xxx_messageInfo_PercentCpc.DiscardUnknown(m)
}

var xxx_messageInfo_PercentCpc proto.InternalMessageInfo

func (m *PercentCpc) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *PercentCpc) GetEnhancedCpcEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.EnhancedCpcEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*EnhancedCpc)(nil), "google.ads.googleads.v1.common.EnhancedCpc")
	proto.RegisterType((*ManualCpc)(nil), "google.ads.googleads.v1.common.ManualCpc")
	proto.RegisterType((*ManualCpm)(nil), "google.ads.googleads.v1.common.ManualCpm")
	proto.RegisterType((*ManualCpv)(nil), "google.ads.googleads.v1.common.ManualCpv")
	proto.RegisterType((*MaximizeConversions)(nil), "google.ads.googleads.v1.common.MaximizeConversions")
	proto.RegisterType((*MaximizeConversionValue)(nil), "google.ads.googleads.v1.common.MaximizeConversionValue")
	proto.RegisterType((*PageOnePromoted)(nil), "google.ads.googleads.v1.common.PageOnePromoted")
	proto.RegisterType((*TargetCpa)(nil), "google.ads.googleads.v1.common.TargetCpa")
	proto.RegisterType((*TargetCpm)(nil), "google.ads.googleads.v1.common.TargetCpm")
	proto.RegisterType((*TargetImpressionShare)(nil), "google.ads.googleads.v1.common.TargetImpressionShare")
	proto.RegisterType((*TargetOutrankShare)(nil), "google.ads.googleads.v1.common.TargetOutrankShare")
	proto.RegisterType((*TargetRoas)(nil), "google.ads.googleads.v1.common.TargetRoas")
	proto.RegisterType((*TargetSpend)(nil), "google.ads.googleads.v1.common.TargetSpend")
	proto.RegisterType((*PercentCpc)(nil), "google.ads.googleads.v1.common.PercentCpc")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/bidding.proto", fileDescriptor_bidding_ebc1284101877700)
}

var fileDescriptor_bidding_ebc1284101877700 = []byte{
	// 927 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xc0, 0xb5, 0x76, 0x5b, 0xd1, 0x71, 0x4a, 0xc9, 0xa6, 0x2d, 0x56, 0x5a, 0x55, 0xd5, 0x4a,
	0x08, 0x0e, 0x68, 0x57, 0x4e, 0x11, 0x07, 0x23, 0x84, 0xec, 0x4d, 0x1a, 0x59, 0x24, 0xaa, 0xb1,
	0x4b, 0x80, 0x28, 0x62, 0x34, 0x9e, 0x99, 0x6c, 0x46, 0xdd, 0x9d, 0xb7, 0xcc, 0xec, 0x26, 0xa4,
	0x17, 0xbe, 0x0b, 0x47, 0xc4, 0x47, 0xe0, 0x8e, 0x84, 0xf8, 0x1e, 0x48, 0x70, 0xe5, 0x03, 0xa0,
	0x9d, 0x99, 0xb5, 0x5b, 0x9c, 0xc4, 0x4d, 0x30, 0x52, 0x4f, 0x9e, 0xf1, 0xbc, 0xf7, 0x7b, 0xff,
	0x66, 0xde, 0x3e, 0xf4, 0x61, 0x02, 0x90, 0xa4, 0x3c, 0x22, 0x4c, 0x47, 0x76, 0x59, 0xad, 0x8e,
	0x3b, 0x11, 0x85, 0x2c, 0x03, 0x19, 0x4d, 0x04, 0x63, 0x42, 0x26, 0x61, 0xae, 0xa0, 0x00, 0xff,
	0xa1, 0x15, 0x09, 0x09, 0xd3, 0xe1, 0x54, 0x3a, 0x3c, 0xee, 0x84, 0x56, 0x7a, 0x3d, 0x3e, 0x8f,
	0xc6, 0x65, 0x99, 0xe9, 0x28, 0x27, 0x09, 0xc7, 0x20, 0x39, 0xce, 0x15, 0x64, 0x50, 0x70, 0x86,
	0x75, 0xa1, 0x48, 0xc1, 0x93, 0x53, 0x9c, 0x00, 0x49, 0xad, 0x91, 0xf5, 0xcd, 0x8b, 0x21, 0x05,
	0x51, 0x09, 0x2f, 0xb0, 0xc8, 0x72, 0xc5, 0xb5, 0x16, 0x20, 0xb1, 0x3e, 0x22, 0x8a, 0xe3, 0x14,
	0x28, 0x29, 0x04, 0x48, 0x47, 0x71, 0xae, 0x46, 0x66, 0x37, 0x29, 0x0f, 0xa3, 0x13, 0x45, 0xf2,
	0x9c, 0x2b, 0xed, 0xce, 0x1f, 0xd4, 0x56, 0x72, 0x11, 0x11, 0x29, 0xa1, 0x30, 0xca, 0xee, 0x34,
	0xb8, 0x85, 0x5a, 0x5b, 0xf2, 0x88, 0x48, 0xca, 0x59, 0x9c, 0xd3, 0xe0, 0x1b, 0x74, 0x73, 0x97,
	0xc8, 0x92, 0xa4, 0x71, 0x4e, 0xfd, 0x1d, 0x74, 0x87, 0xbb, 0x33, 0x4c, 0x73, 0x8a, 0xb9, 0x24,
	0x93, 0x94, 0xb3, 0xb6, 0xf7, 0xc8, 0xfb, 0xa0, 0xb5, 0xb1, 0xee, 0x12, 0x13, 0xd6, 0x86, 0xc3,
	0x3e, 0x40, 0xba, 0x47, 0xd2, 0x92, 0x8f, 0x7c, 0x3e, 0x63, 0x6e, 0x59, 0xad, 0xa0, 0x35, 0x43,
	0x67, 0x2f, 0x6f, 0x8e, 0x83, 0xbb, 0x68, 0x6d, 0x97, 0x7c, 0x2f, 0x32, 0xf1, 0x82, 0xc7, 0x20,
	0x8f, 0xb9, 0xaa, 0x82, 0xd5, 0xc1, 0xd7, 0xe8, 0xdd, 0xf9, 0xbf, 0x0d, 0xdf, 0xff, 0x14, 0xb5,
	0x5c, 0x76, 0x14, 0x10, 0xed, 0x1c, 0x7a, 0x30, 0xe7, 0xd0, 0x26, 0x94, 0x93, 0x94, 0x5b, 0x97,
	0x90, 0x55, 0x18, 0x01, 0xd1, 0xc1, 0xef, 0xd7, 0xd0, 0xed, 0x21, 0x49, 0xf8, 0x53, 0xc9, 0x87,
	0xae, 0x40, 0xfe, 0x0f, 0xe8, 0xd6, 0x2b, 0x35, 0x32, 0xd0, 0xb7, 0x37, 0xf6, 0xc3, 0xf3, 0x6e,
	0x82, 0x29, 0x52, 0xf8, 0x2f, 0xcc, 0xd8, 0x21, 0xb6, 0x81, 0xa4, 0x5b, 0xb2, 0xcc, 0x2e, 0x3a,
	0x1f, 0xad, 0xe8, 0x97, 0x76, 0xfe, 0x10, 0xdd, 0xab, 0x92, 0x3c, 0x11, 0x0c, 0x53, 0x2e, 0x52,
	0x21, 0x13, 0x9c, 0x09, 0xaa, 0x40, 0xb7, 0x1b, 0x26, 0xbc, 0xfb, 0x73, 0xe1, 0x0d, 0x64, 0xf1,
	0xf1, 0x47, 0x36, 0xba, 0x35, 0x9a, 0xd3, 0xbe, 0x60, 0xb1, 0x55, 0xdc, 0x35, 0x7a, 0xfe, 0x67,
	0x68, 0xa5, 0xa2, 0x65, 0xc0, 0xc4, 0xa1, 0xe0, 0xaa, 0xdd, 0x7c, 0x8d, 0x34, 0xb5, 0x26, 0x82,
	0xed, 0x3a, 0x05, 0x7f, 0x80, 0xd6, 0x40, 0xa6, 0xa7, 0x58, 0x11, 0xa1, 0x39, 0x76, 0xde, 0xe9,
	0xf6, 0xb5, 0x85, 0xf5, 0x7f, 0xa7, 0x52, 0x1b, 0x55, 0x5a, 0xb1, 0xf1, 0x4b, 0xfb, 0xcf, 0xd1,
	0x7b, 0xaf, 0x50, 0xf0, 0xc9, 0x11, 0x97, 0x78, 0x52, 0xb2, 0xaa, 0x88, 0x14, 0x64, 0x95, 0x0a,
	0x21, 0x39, 0x6b, 0x5f, 0x5f, 0x08, 0x7f, 0xa4, 0x66, 0xe0, 0xaf, 0x8e, 0xb8, 0xec, 0x1b, 0x48,
	0x3c, 0x63, 0xf8, 0x19, 0x7a, 0xff, 0x0c, 0x63, 0xdf, 0x95, 0x24, 0x15, 0xc5, 0x29, 0xd6, 0x14,
	0x14, 0xc7, 0x42, 0xe3, 0x14, 0x4e, 0xda, 0x37, 0x2e, 0x6d, 0xee, 0x0b, 0x8b, 0x19, 0x57, 0x94,
	0x81, 0xde, 0x81, 0x93, 0xe0, 0x6f, 0x0f, 0xdd, 0x7c, 0x66, 0x6e, 0x57, 0x9c, 0x13, 0x7f, 0x1b,
	0xad, 0xba, 0xbb, 0x49, 0x73, 0x52, 0x97, 0xd0, 0x5b, 0x5c, 0xc2, 0xdb, 0x45, 0x8d, 0x70, 0xe5,
	0x5b, 0xfe, 0x85, 0xd8, 0x41, 0x77, 0x6a, 0xe2, 0x61, 0x0a, 0xa0, 0x6a, 0x5e, 0x73, 0x31, 0x6f,
	0xd5, 0xf2, 0x9e, 0x54, 0x6a, 0x96, 0x56, 0xbd, 0xe1, 0x3a, 0xea, 0x2c, 0xf8, 0xa5, 0x81, 0xee,
	0xda, 0xdd, 0x60, 0xda, 0xaf, 0xc6, 0x55, 0xbb, 0xf2, 0x5f, 0xa0, 0xb7, 0xea, 0x8e, 0xe5, 0xde,
	0xd4, 0xb7, 0x0b, 0xde, 0xd4, 0x99, 0x9c, 0x1d, 0xc7, 0x30, 0xaf, 0xea, 0x42, 0x89, 0xd1, 0xd4,
	0x9e, 0xff, 0x25, 0x6a, 0xd7, 0x6b, 0x7c, 0xa8, 0x08, 0x35, 0x8b, 0xd7, 0x4f, 0xe2, 0xbd, 0x5a,
	0xf9, 0x89, 0xd3, 0x5d, 0x58, 0x99, 0xe6, 0xd5, 0x2a, 0x13, 0xfc, 0xda, 0x44, 0xbe, 0x0d, 0xea,
	0x69, 0x59, 0x28, 0x22, 0x9f, 0xdb, 0xdc, 0xed, 0xa3, 0xfb, 0xee, 0x2e, 0x81, 0xfd, 0xdb, 0x7d,
	0x02, 0x16, 0xdf, 0xaa, 0xc7, 0x1b, 0xd6, 0x5a, 0xbb, 0x98, 0xa3, 0xba, 0x20, 0x06, 0x68, 0x95,
	0x42, 0x96, 0xf3, 0x42, 0x14, 0xa0, 0x30, 0x83, 0x8c, 0x08, 0xe9, 0x92, 0x32, 0xdf, 0x22, 0xc6,
	0x85, 0x12, 0x32, 0x71, 0x8f, 0x7b, 0xa6, 0xb6, 0x69, 0xb4, 0x96, 0x9f, 0x8f, 0x65, 0x76, 0x9e,
	0x4b, 0x34, 0x83, 0xeb, 0x4b, 0x68, 0x06, 0x7f, 0x79, 0x08, 0x3d, 0x9b, 0x7e, 0x6a, 0xfe, 0xe3,
	0x97, 0xea, 0x8d, 0xef, 0x01, 0x3f, 0x7b, 0xa8, 0x65, 0xa3, 0x1d, 0xe7, 0x5c, 0x32, 0xff, 0x73,
	0xb4, 0xe6, 0xc2, 0xd5, 0xd5, 0xfe, 0x12, 0xed, 0xcf, 0x35, 0x4d, 0x83, 0xf9, 0xbf, 0x1a, 0x60,
	0xe5, 0x2e, 0x1a, 0x72, 0x45, 0xb9, 0x2c, 0xaa, 0x01, 0xe7, 0x7c, 0x03, 0xde, 0xd5, 0xb3, 0x7b,
	0xe6, 0xc8, 0xd4, 0xb8, 0xca, 0xc8, 0xd4, 0xff, 0xc3, 0x43, 0x01, 0x85, 0x2c, 0xbc, 0x78, 0x18,
	0xed, 0xaf, 0xf4, 0xed, 0xec, 0x3a, 0xac, 0xa8, 0x43, 0x6f, 0xdf, 0xcd, 0x95, 0x61, 0x02, 0x29,
	0x91, 0x49, 0x08, 0x2a, 0x89, 0x12, 0x2e, 0x8d, 0xcd, 0x7a, 0xce, 0xcc, 0x85, 0x3e, 0x6f, 0x12,
	0xfe, 0xc4, 0xfe, 0xfc, 0xd8, 0x68, 0x6e, 0xf7, 0x7a, 0x3f, 0x35, 0x1e, 0x6e, 0x5b, 0x58, 0x8f,
	0xe9, 0xd0, 0x2e, 0xab, 0xd5, 0x5e, 0x27, 0x8c, 0x8d, 0xd8, 0x6f, 0xb5, 0xc0, 0x41, 0x8f, 0xe9,
	0x83, 0xa9, 0xc0, 0xc1, 0x5e, 0xe7, 0xc0, 0x0a, 0xfc, 0xd9, 0x08, 0xec, 0xbf, 0xdd, 0x6e, 0x8f,
	0xe9, 0x6e, 0x77, 0x2a, 0xd2, 0xed, 0xee, 0x75, 0xba, 0x5d, 0x2b, 0x34, 0xb9, 0x61, 0xbc, 0x7b,
	0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x64, 0x1a, 0x67, 0xa6, 0x0b, 0x00, 0x00,
}
