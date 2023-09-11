/*
Payment:

	description: Describes a payment
	type: object
	properties:
	  uri:
	    type: string
	    description: 'A payment uri to be called by the Buyer App. If empty, then the payment is to be done offline. The details of payment should be present in the params object. If ```tl_method``` = http/get, then the payment details will be sent as url params. Two url param values, ```$transaction_id``` and ```$amount``` are mandatory. And example url would be : https://www.example.com/pay?txid=$transaction_id&amount=$amount&vpa=upiid&payee=shopez&billno=1234'
	    format: uri
	  tl_method:
	    type: string
	    enum:
	      - http/get
	      - http/post
	      - payto
	      - upi
	  params:
	    type: object
	    properties:
	      transaction_id:
	        type: string
	        description: This value will be placed in the the $transaction_id url param in case of http/get and in the requestBody http/post requests
	      transaction_status:
	        type: string
	      amount:
	        $ref: '#/components/schemas/Price/properties/value'
	      currency:
	        $ref: '#/components/schemas/Price/properties/currency'
	    additionalProperties:
	      type: string
	    required:
	      - currency
	  type:
	    type: string
	    enum:
	      - ON-ORDER
	      - PRE-FULFILLMENT
	      - ON-FULFILLMENT
	      - POST-FULFILLMENT
	  status:
	    type: string
	    enum:
	      - PAID
	      - NOT-PAID
	  time:
	    $ref: '#/components/schemas/Time'
	  collected_by:
	    type: string
	    enum:
	      - BAP
	      - BPP
	  "@ondc/org/collected_by_status":
	    type: string
	    enum:
	      - Assert
	      - Agree
	      - Disagree
	      - Terminate
	  "@ondc/org/buyer_app_finder_fee_type":
	    type: string
	    enum:
	      - Amount
	      - Percent
	  "@ondc/org/buyer_app_finder_fee_amount":
	    $ref: '#/components/schemas/DecimalValue'
	  "@ondc/org/withholding_amount":
	    $ref: '#/components/schemas/DecimalValue'
	  "@ondc/org/withholding_amount_status":
	    type: string
	    enum:
	      - Assert
	      - Agree
	      - Disagree
	      - Terminate
	  "@ondc/org/return_window":
	    description: return window for withholding amount in ISO8601 durations format e.g. 'PT24H' indicates 24 hour return window
	    type: string
	  "@ondc/org/return_window_status":
	    type: string
	    enum:
	      - Assert
	      - Agree
	      - Disagree
	      - Terminate
	  "@ondc/org/settlement_basis":
	    description: In case of prepaid payment, whether settlement between counterparties should be on the basis of collection, shipment or delivery
	    type: string
	    enum:
	      - shipment
	      - delivery
	      - return_window_expiry
	  "@ondc/org/settlement_basis_status":
	    type: string
	    enum:
	      - Assert
	      - Agree
	      - Disagree
	      - Terminate
	  "@ondc/org/settlement_window":
	    description: settlement window for the counterparty in ISO8601 durations format e.g. 'PT48H' indicates 48 hour return window
	    type: string
	  "@ondc/org/settlement_window_status":
	    type: string
	    enum:
	      - Assert
	      - Agree
	      - Disagree
	      - Terminate
	  "@ondc/org/settlement_details":
	    type: array
	    items:
	      type: object
	      properties:
	        settlement_counterparty:
	          type: string
	          enum:
	            - buyer
	            - buyer-app
	            - seller-app
	            - logistics-provider
	        settlement_phase:
	          type: string
	          enum:
	            - sale-amount
	            - withholding-amount
	            - refund
	        settlement_amount:
	          type: integer
	        settlement_type:
	          type: string
	          enum:
	            - neft
	            - rtgs
	            - upi
	            - credit
	        settlement_bank_account_no:
	          type: string
	        settlement_ifsc_code:
	          type: string
	        upi_address:
	          description: UPI payment address e.g. VPA
	          type: string
	        bank_name:
	          description: Bank name
	          type: string
	        branch_name:
	          description: Branch name
	          type: string
	        beneficiary_name:
	          description: Beneficiary Name
	          type: string
	        beneficiary_address:
	          description: Beneficiary Address
	          type: string
	        settlement_status:
	          type: string
	          enum:
	            - PAID
	            - NOT-PAID
	        settlement_reference:
	          description: Settlement transaction reference number
	          type: string
	        settlement_timestamp:
	          description: Settlement transaction timestamp
	          type: string
	          format: date-time
*/
package schema

import "encoding/json"

type PaymentParams struct {
	TransactionId        string            `json:"transaction_id,omitempty"`     // This value will be placed in the the $transaction_id url param in case of http/get and in the requestBody http/post requests
	TransactionStatus    string            `json:"transaction_status,omitempty"` //
	Amount               PriceValue        `json:"amount,omitempty"`             //
	Currency             PriceCurrency     `json:"currency,omitempty"`           //
	AdditionalProperties map[string]string `json:"-"`                            // Additional properties
}

// custom marshaler
func (p PaymentParams) MarshalJSON() ([]byte, error) {
	//flatten the additional properties in the json as kv pairs

	var result = make(map[string]interface{})

	result["transaction_id"] = p.TransactionId
	result["transaction_status"] = p.TransactionStatus
	result["amount"] = p.Amount
	result["currency"] = p.Currency

	for k, v := range p.AdditionalProperties {
		result[k] = v
	}

	return json.Marshal(result)
}

// custom unmarshaler
func (p *PaymentParams) UnmarshalJSON(data []byte) error {
	//unflatten the additional properties in the json as kv pairs

	var result = make(map[string]interface{})

	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	for k, v := range result {
		switch k {
		case "transaction_id":
			p.TransactionId = v.(string)
		case "transaction_status":
			p.TransactionStatus = v.(string)
		case "amount":
			p.Amount = v.(PriceValue)
		case "currency":
			p.Currency = v.(PriceCurrency)
		default:
			p.AdditionalProperties[k] = v.(string)
		}
	}

	return nil
}

type PaymentTlMethodType string

const (
	PaymentTlMethodTypeHttpGet  PaymentTlMethodType = "http/get"
	PaymentTlMethodTypeHttpPost PaymentTlMethodType = "http/post"
	PaymentTlMethodTypePayto    PaymentTlMethodType = "payto"
	PaymentTlMethodTypeUpi      PaymentTlMethodType = "upi"
)

type PaymentType string

const (
	PaymentTypeOnOrder         PaymentType = "ON-ORDER"
	PaymentTypePreFulfillment  PaymentType = "PRE-FULFILLMENT"
	PaymentTypeOnFulfillment   PaymentType = "ON-FULFILLMENT"
	PaymentTypePostFulfillment PaymentType = "POST-FULFILLMENT"
)

type SettlementStatus string

const (
	SettlementStatusPaid    SettlementStatus = "PAID"
	SettlementStatusNotPaid SettlementStatus = "NOT-PAID"
)

type PaymentCollectedBy string

const (
	PaymentCollectedByBap PaymentCollectedBy = "BAP"
	PaymentCollectedByBpp PaymentCollectedBy = "BPP"
)

type AgreementStatus string

const (
	AgreementStatusAssert    AgreementStatus = "Assert"
	AgreementStatusAgree     AgreementStatus = "Agree"
	AgreementStatusDisagree  AgreementStatus = "Disagree"
	AgreementStatusTerminate AgreementStatus = "Terminate"
)

type PaymentBuyerAppFinderFeeType string

const (
	PaymentBuyerAppFinderFeeTypeAmount  PaymentBuyerAppFinderFeeType = "Amount"
	PaymentBuyerAppFinderFeeTypePercent PaymentBuyerAppFinderFeeType = "Percent"
)

type SettlementCounterParty string

const (
	SettlementCounterPartyBuyer             SettlementCounterParty = "buyer"
	SettlementCounterPartyBuyerApp          SettlementCounterParty = "buyer-app"
	SettlementCounterPartySellerApp         SettlementCounterParty = "seller-app"
	SettlementCounterPartyLogisticsProvider SettlementCounterParty = "logistics-provider"
)

type SettlementPhase string

const (
	SettlementPhaseSaleAmount        SettlementPhase = "sale-amount"
	SettlementPhaseWithholdingAmount SettlementPhase = "withholding-amount"
	SettlementPhaseRefund            SettlementPhase = "refund"
)

type SettlementBasis string

const (
	SettlementBasisShipment           SettlementBasis = "shipment"
	SettlementBasisDelivery           SettlementBasis = "delivery"
	SettlementBasisReturnWindowExpiry SettlementBasis = "return_window_expiry"
)

type SettlementType string

const (
	SettlementTypeNeft   SettlementType = "neft"
	SettlementTypeRtgs   SettlementType = "rtgs"
	SettlementTypeUpi    SettlementType = "upi"
	SettlementTypeCredit SettlementType = "credit"
)

type PaymentSettlementDetails struct {
	SettlementCounterParty  SettlementCounterParty `json:"settlement_counterparty,omitempty"`
	SettlementPhase         SettlementPhase        `json:"settlement_phase,omitempty"`
	SettlementAmount        int                    `json:"settlement_amount,omitempty"`
	SettlementType          SettlementType         `json:"settlement_type,omitempty"`
	SettlementBankAccountNo string                 `json:"settlement_bank_account_no,omitempty"`
	SettlementIfscCode      string                 `json:"settlement_ifsc_code,omitempty"`
	UpiAddress              string                 `json:"upi_address,omitempty"`
	BankName                string                 `json:"bank_name,omitempty"`
	BranchName              string                 `json:"branch_name,omitempty"`
	BeneficiaryName         string                 `json:"beneficiary_name,omitempty"`
	BeneficiaryAddress      string                 `json:"beneficiary_address,omitempty"`
	SettlementStatus        SettlementStatus       `json:"settlement_status,omitempty"`
	SettlementReference     string                 `json:"settlement_reference,omitempty"`
	SettlementTimestamp     string                 `json:"settlement_timestamp,omitempty"`
}

type Payment struct {
	Uri                            string                       `json:"uri,omitempty"`                           // A payment uri to be called by the Buyer App
	TlMethod                       PaymentTlMethodType          `json:"tl_method,omitempty"`                     // If empty, then the payment is to be done offline
	Params                         PaymentParams                `json:"params,omitempty"`                        // The details of payment should be present in the params object
	Type                           PaymentType                  `json:"type,omitempty"`                          //
	Status                         string                       `json:"status,omitempty"`                        //
	Time                           Time                         `json:"time,omitempty"`                          //
	CollectedBy                    PaymentCollectedBy           `json:"collected_by,omitempty"`                  //
	OndcOrgCollectedByStatus       AgreementStatus              `json:"@ondc/org/collected_by_status,omitempty"` //
	OndcOrgBuyerAppFinderFeeType   PaymentBuyerAppFinderFeeType `json:"@ondc/org/buyer_app_finder_fee_type,omitempty"`
	OndcOrgBuyerAppFinderFeeAmount DecimalValue                 `json:"@ondc/org/buyer_app_finder_fee_amount,omitempty"`
	OndcOrgWithholdingAmount       DecimalValue                 `json:"@ondc/org/withholding_amount,omitempty"`
	OndcOrgWithholdingAmountStatus AgreementStatus              `json:"@ondc/org/withholding_amount_status,omitempty"`
	OndcOrgReturnWindow            string                       `json:"@ondc/org/return_window,omitempty"`
	OndcOrgReturnWindowStatus      AgreementStatus              `json:"@ondc/org/return_window_status,omitempty"`
	OndcOrgSettlementBasis         SettlementBasis              `json:"@ondc/org/settlement_basis,omitempty"`
	OndcOrgSettlementBasisStatus   AgreementStatus              `json:"@ondc/org/settlement_basis_status,omitempty"`
	OndcOrgSettlementWindow        string                       `json:"@ondc/org/settlement_window,omitempty"`
	OndcOrgSettlementWindowStatus  AgreementStatus              `json:"@ondc/org/settlement_window_status,omitempty"`
	OndcOrgSettlementDetails       []PaymentSettlementDetails   `json:"@ondc/org/settlement_details,omitempty"`
}
