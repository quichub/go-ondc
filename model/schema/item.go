package schema

type ItemId string

/*


   Item:
     description: Describes a product or a service offered to the end consumer by the provider
     type: object
     properties:
       id:
         description: This is the most unique identifier of a service item. An example of an Item ID could be the SKU of a product.
         type: string
       parent_item_id:
         $ref: '#/components/schemas/Item/properties/id'
       descriptor:
         $ref: '#/components/schemas/Descriptor'
       price:
         $ref: '#/components/schemas/Price'
       quantity:
         $ref: '#/components/schemas/ItemQuantity'
       category_id:
         $ref: '#/components/schemas/Category/properties/id'
       category_ids:
         description: Categories this item can be listed under
         type: array
         items:
           allOf:
             - $ref: '#/components/schemas/Category/properties/id'
       fulfillment_id:
         $ref: '#/components/schemas/Fulfillment/properties/id'
       rating:
         $ref: '#/components/schemas/Rating/properties/value'
       location_id:
         $ref: '#/components/schemas/Location/properties/id'
       time:
         $ref: '#/components/schemas/Time'
       rateable:
         $ref: '#/components/schemas/Rateable'
       matched:
         type: boolean
       related:
         type: boolean
       recommended:
         type: boolean
       "@ondc/org/returnable":
         description: whether the item is returnable
         type: boolean
       "@ondc/org/seller_pickup_return":
         description: in case of return, whether the item should be picked up by seller
         type: boolean
       "@ondc/org/return_window":
         description: return window for the item in ISO8601 durations format e.g. 'PT24H' indicates 24 hour return window. Mandatory if "@ondc/org/returnable" is "true"
         type: string
       "@ondc/org/cancellable":
         description: whether the item is cancellable
         type: boolean
       "@ondc/org/time_to_ship":
         description: time from order confirmation by which item ready to ship in ISO8601 durations format (e.g. 'PT30M' indicates item ready to ship in 30 mins). Mandatory for category_id "F&B"
         type: string
       "@ondc/org/available_on_cod":
         description: whether the catalog item is available on COD
         type: boolean
       "@ondc/org/contact_details_consumer_care":
         description: contact details for consumer care
         type: string
       "@ondc/org/statutory_reqs_packaged_commodities":
         type: object
         description:
           mandatory for category_id "Packaged Commodities"
         properties:
           manufacturer_or_packer_name:
             description: name of manufacturer or packer (in case manufacturer is not the packer) or name of importer for imported goods
             type: string
           manufacturer_or_packer_address:
             description: address of manufacturer or packer (in case manufacturer is not the packer) or address of importer for imported goods
             type: string
           mfg_license_no:
             description: manufacturing license no
             type: string
           common_or_generic_name_of_commodity:
             description: common or generic name of commodity
             type: string
           multiple_products_name_number_or_qty:
             description: for packages with multiple products, the name and number of quantity of each (can be shown as "name1-number_or_quantity; name2-number_or_quantity..")
             type: string
           net_quantity_or_measure_of_commodity_in_pkg:
             description: net quantity of commodity in terms of standard unit of weight or measure of commodity contained in package
             type: string
           month_year_of_manufacture_packing_import:
             description: month and year of manufacture or packing or import
             type: string
           expiry_date:
             description: month and year of expiry
             type: string
       "@ondc/org/statutory_reqs_prepackaged_food":
         type: object
         description:
           mandatory for category_id "Packaged food"
         properties:
           ingredients_info:
             description: list of ingredients (except single ingredient foods), can be shown as ingredient (with percentage); ingredient (with percentage);..)
                          e.g. "Puffed Rice (40%); Split Green Gram (20%); Ground Nuts (20%);.."
             type: string
           nutritional_info:
             description: nutritional info (can be shown as nutritional info (with unit, per standard unit, per serving);..)
                          e.g. "Energy(KCal) - (per 100kg) 420, (per serving 50g) 250; Protein(g) - (per 100kg) 12, (per serving 50g)6;.."
             type: string
           additives_info:
             description: food additives together with specific name or recognized International Numbering System (can be shown as additive1-name or number;additive2-name or number;..)
             type: string
           manufacturer_or_packer_name:
             description: name of manufacturer or packer (for non-retail containers)
             type: string
           manufacturer_or_packer_address:
             description: address of manufacturer or packer (for non-retail containers)
             type: string
           brand_owner_name:
             description: name of brand owner
             type: string
           brand_owner_address:
             description: address of brand owner
             type: string
           brand_owner_FSSAI_logo:
             description: FSSAI logo of brand owner (url based image e.g. uri:http://path/to/image)
             type: string
           brand_owner_FSSAI_license_no:
             description: FSSAI license no of brand owner
             type: string
           other_FSSAI_license_no:
             description: FSSAI license no of manufacturer or marketer or packer or bottler if different from brand owner
             type: string
           net_quantity:
             description: net quantity
             type: string
           importer_name:
             description: name of importer
             type: string
           importer_address:
             description: address of importer
             type: string
           importer_FSSAI_logo:
             description: FSSAI logo of importer (url based image e.g. uri:http://path/to/image)
             type: string
           importer_FSSAI_license_no:
             description: FSSAI license no of importer
             type: string
           imported_product_country_of_origin:
             description: country of origin for imported products (ISO 3166 Alpha-3 code format)
             type: string
           other_importer_name:
             description: name of importer for product manufactured outside but packaged or bottled in India
             type: string
           other_importer_address:
             description: address of importer for product manufactured outside but packaged or bottled in India
             type: string
           other_premises:
             description: premises where product manufactured outside are packaged or bottled in India
             type: string
       tags:
         $ref: '#/components/schemas/TagGroup'
*/

type PackagedCommodity struct {
	ManufacturerOrPackerName             string `json:"manufacturer_or_packer_name,omitempty"`                 // name of manufacturer or packer (in case manufacturer is not the packer) or name of importer for imported goods
	ManufacturerOrPackerAddress          string `json:"manufacturer_or_packer_address,omitempty"`              // address of manufacturer or packer (in case manufacturer is not the packer) or address of importer for imported goods
	MfgLicenseNo                         string `json:"mfg_license_no,omitempty"`                              // manufacturing license no
	CommonOrGenericNameOfCommodity       string `json:"common_or_generic_name_of_commodity,omitempty"`         // common or generic name of commodity
	MultipleProductsNameNumberOrQty      string `json:"multiple_products_name_number_or_qty,omitempty"`        // for packages with multiple products, the name and number of quantity of each (can be shown as "name1-number_or_quantity; name2-number_or_quantity..")
	NetQuantityOrMeasureOfCommodityInPkg string `json:"net_quantity_or_measure_of_commodity_in_pkg,omitempty"` // net quantity of commodity in terms of standard unit of weight or measure of commodity contained in package
	MonthYearOfManufacturePackingImport  string `json:"month_year_of_manufacture_packing_import,omitempty"`    // month and year of manufacture or packing or import
	ExpiryDate                           string `json:"expiry_date,omitempty"`                                 // month and year of expiry
}

type PackagedFood struct {
	IngredientsInfo                string `json:"ingredients_info,omitempty"`                   // list of ingredients (except single ingredient foods), can be shown as ingredient (with percentage); ingredient (with percentage);..)
	NutritionalInfo                string `json:"nutritional_info,omitempty"`                   // nutritional info (can be shown as nutritional info (with unit, per standard unit, per serving);..)
	AdditivesInfo                  string `json:"additives_info,omitempty"`                     // food additives together with specific name or recognized International Numbering System (can be shown as additive1-name or number;additive2-name or number;..)
	ManufacturerOrPackerName       string `json:"manufacturer_or_packer_name,omitempty"`        // name of manufacturer or packer (for non-retail containers)
	ManufacturerOrPackerAddress    string `json:"manufacturer_or_packer_address,omitempty"`     // address of manufacturer or packer (for non-retail containers)
	BrandOwnerName                 string `json:"brand_owner_name,omitempty"`                   // name of brand owner
	BrandOwnerAddress              string `json:"brand_owner_address,omitempty"`                // address of brand owner
	BrandOwnerFSSAILogo            string `json:"brand_owner_FSSAI_logo,omitempty"`             // FSSAI logo of brand owner (url based image e.g. uri:http://path/to/image)
	BrandOwnerFSSAILicenseNo       string `json:"brand_owner_FSSAI_license_no,omitempty"`       // FSSAI license no of brand owner
	OtherFSSAILicenseNo            string `json:"other_FSSAI_license_no,omitempty"`             // FSSAI license no of manufacturer or marketer or packer or bottler if different from brand owner
	NetQuantity                    string `json:"net_quantity,omitempty"`                       // net quantity
	ImporterName                   string `json:"importer_name,omitempty"`                      // name of importer
	ImporterAddress                string `json:"importer_address,omitempty"`                   // address of importer
	ImporterFSSAILogo              string `json:"importer_FSSAI_logo,omitempty"`                // FSSAI logo of importer (url based image e.g. uri:http://path/to/image)
	ImporterFSSAILicenseNo         string `json:"importer_FSSAI_license_no,omitempty"`          // FSSAI license no of importer
	ImportedProductCountryOfOrigin string `json:"imported_product_country_of_origin,omitempty"` // country of origin for imported products (ISO 3166 Alpha-3 code format)
	OtherImporterName              string `json:"other_importer_name,omitempty"`                // name of importer for product manufactured outside but packaged or bottled in India
	OtherImporterAddress           string `json:"other_importer_address,omitempty"`             // address of importer for product manufactured outside but packaged or bottled in India
	OtherPremises                  string `json:"other_premises,omitempty"`                     // premises where product manufactured outside are packaged or bottled in India
}

type Item struct {
	Id                               ItemId             `json:"id,omitempty"` // This is the most unique identifier of a service item. An example of an Item ID could be the SKU of a product.
	ParentItemId                     *ItemId            `json:"parent_item_id,omitempty"`
	Descriptor                       Descriptor         `json:"descriptor,omitempty"`
	Price                            Price              `json:"price,omitempty"`
	Quantity                         ItemQuantity       `json:"quantity,omitempty"`
	CategoryId                       *CategoryId        `json:"category_id,omitempty"`
	CategoryIds                      []CategoryId       `json:"category_ids,omitempty"` // Categories this item can be listed under
	FulfillmentId                    *FulfillmentId     `json:"fulfillment_id,omitempty"`
	Rating                           *RatingId          `json:"rating,omitempty"`
	LocationId                       *LocationId        `json:"location_id,omitempty"`
	Time                             *Time              `json:"time,omitempty"`
	Rateable                         *Rateable          `json:"rateable,omitempty"`
	Matched                          *bool              `json:"matched,omitempty"`
	Related                          *bool              `json:"related,omitempty"`
	Recommended                      *bool              `json:"recommended,omitempty"`
	Returnable                       *bool              `json:"@ondc/org/returnable,omitempty"`                    // whether the item is returnable
	SellerPickupReturn               *bool              `json:"@ondc/org/seller_pickup_return,omitempty"`          // in case of return, whether the item should be picked up by seller
	ReturnWindow                     *string            `json:"@ondc/org/return_window,omitempty"`                 // return window for the item in ISO8601 durations format e.g. 'PT24H' indicates 24 hour return window. Mandatory if "@ondc/org/returnable" is "true"
	Cancellable                      *bool              `json:"@ondc/org/cancellable,omitempty"`                   // whether the item is cancellable
	TimeToShip                       *string            `json:"@ondc/org/time_to_ship,omitempty"`                  // time from order confirmation by which item ready to ship in ISO8601 durations format (e.g. 'PT30M' indicates item ready to ship in 30 mins). Mandatory for category_id "F&B"
	AvailableOnCod                   *bool              `json:"@ondc/org/available_on_cod,omitempty"`              // whether the catalog item is available on COD
	ContactDetailsConsumerCare       *string            `json:"@ondc/org/contact_details_consumer_care,omitempty"` // contact details for consumer care
	StatutoryReqsPackagedCommodities *PackagedCommodity `json:"@ondc/org/statutory_reqs_packaged_commodities,omitempty"`
	StatutoryReqsPrepackagedFood     *PackagedFood      `json:"@ondc/org/statutory_reqs_prepackaged_food,omitempty"`
	Tags                             *TagGroup          `json:"tags,omitempty"`
}
