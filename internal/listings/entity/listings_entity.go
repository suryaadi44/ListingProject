package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Listing struct {
	ListingId             string               `bson:"_id"`
	Url                   string               `bson:"listing_url"`
	Images                ImagesDetail         `bson:"images"`
	Name                  string               `bson:"name"`
	Summary               string               `bson:"summary"`
	Space                 string               `bson:"space"`
	Description           string               `bson:"description"`
	NeighbourhoodOverview string               `bson:"neighborhood_overview"`
	Notes                 string               `bson:"notes"`
	Transit               string               `bson:"transit"`
	Access                string               `bson:"access"`
	Address               AddressDetail        `bson:"address"`
	Interaction           string               `bson:"interaction"`
	HouseRules            string               `bson:"house_rules"`
	PropertyType          string               `bson:"property_type"`
	RoomType              string               `bson:"room_type"`
	BedType               string               `bson:"bed_type"`
	MinNights             string               `bson:"minimum_nights"`
	MaxNights             string               `bson:"maximum_nights"`
	CancellationPolicy    string               `bson:"cancellation_policy"`
	Price                 primitive.Decimal128 `bson:"price"`
	WeeklyPrice           primitive.Decimal128 `bson:"weekly_price"`
	MonthlyPrice          primitive.Decimal128 `bson:"monthly_price"`
	Deposit               primitive.Decimal128 `bson:"security_deposit"`
	CleaningFee           primitive.Decimal128 `bson:"cleaning_fee"`
	Accommodates          int64                `bson:"accommodates"`
	Bedrooms              int64                `bson:"bedrooms"`
	Beds                  int64                `bson:"beds"`
	Bathrooms             primitive.Decimal128 `bson:"bathrooms"`
	Amenities             []string             `bson:"amenities"`
	ExtraPeople           primitive.Decimal128 `bson:"extra_people"`
	GuestIncluded         primitive.Decimal128 `bson:"guests_included"`
	Availability          AvailabilityDetail   `bson:"availability"`
	Host                  HostDetail           `bson:"host"`
	NoOfReviews           int64                `bson:"number_of_reviews"`
	ReviewScores          ReviewScoresDetail   `bson:"review_scores"`
	Reviews               []ReviewsDetail      `bson:"reviews"`
	ReviewsPerMonth       int64                `bson:"reviews_per_month"`
	FirstReview           time.Time            `bson:"first_review"`
	LastReview            time.Time            `bson:"last_review"`
	CalendarLastSraped    time.Time            `bson:"calendar_last_scraped"`
	LastSraped            time.Time            `bson:"last_scraped"`
}

type AddressDetail struct {
	Country        string         `bson:"country"`
	CountryCode    string         `bson:"country_code"`
	GovernmentArea string         `bson:"government_area"`
	Location       LocationDetail `bson:"location"`
	Market         string         `bson:"market"`
	Street         string         `bson:"street"`
	Suburb         string         `bson:"suburb"`
}

type LocationDetail struct {
	Type            string    `bson:"type"`
	Coordinates     []float64 `bson:"coordinates"`
	IsLocationExact bool      `bson:"is_location_exact"`
}

type AvailabilityDetail struct {
	Availability30  int64 `bson:"availability_30"`
	Availability60  int64 `bson:"availability_60"`
	Availability90  int64 `bson:"availability_90"`
	Availability365 int64 `bson:"availability_365"`
}

type HostDetail struct {
	HostId              string   `bson:"host_id"`
	Name                string   `bson:"host_name"`
	HostAbout           string   `bson:"host_about"`
	ListingCount        int64    `bson:"host_listings_count"`
	TotalListingCount   int64    `bson:"host_total_listings_count"`
	Location            string   `bson:"host_location"`
	Neighbourhood       string   `bson:"host_neighbourhood"`
	IsHaveProfilePic    bool     `bson:"host_has_profile_pic"`
	PictureUrl          string   `bson:"host_picture_url"`
	ThumbnailUrl        string   `bson:"host_thumbnail_url"`
	IsIdentitiyVerified bool     `bson:"host_identity_verified"`
	IsSuperhost         bool     `bson:"host_is_superhost"`
	ResponseRate        int64    `bson:"host_response_rate"`
	ResponseTime        string   `bson:"host_response_time"`
	Url                 string   `bson:"host_url"`
	Verification        []string `bson:"host_verifications"`
}

type ImagesDetail struct {
	PictureUrl      string `bson:"picture_url"`
	ThumbnailUrl    string `bson:"thumbnail_url"`
	MediumUrl       string `bson:"medium_url"`
	LargePictureUrl string `bson:"xl_picture_url"`
}

type ReviewScoresDetail struct {
	Accuracy      int64 `bson:"review_scores_accuracy"`
	CheckIn       int64 `bson:"review_scores_checkin"`
	Cleanliness   int64 `bson:"review_scores_cleanliness"`
	Communication int64 `bson:"review_scores_communication"`
	Location      int64 `bson:"review_scores_location"`
	Rating        int64 `bson:"review_scores_rating"`
	Value         int64 `bson:"review_scores_value"`
}

type ReviewsDetail struct {
	ReviewsId    string    `bson:"_id"`
	Comments     string    `bson:"comments"`
	Date         time.Time `bson:"date"`
	ListingId    string    `bson:"listing_id"`
	ReviewerId   string    `bson:"reviewer_id"`
	ReviewerName string    `bson:"reviewer_name"`
}

type Listings []Listing
