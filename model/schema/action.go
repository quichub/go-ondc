package schema

/*
action:
          type: string
          description: Defines the ONDC API call. Any actions other than the enumerated actions are not supported by ONDC Protocol
          enum:
            - search
            - select
            - init
            - confirm
            - update
            - status
            - track
            - cancel
            - rating
            - support
            - on_search
            - on_select
            - on_init
            - on_confirm
            - on_update
            - on_status
            - on_track
            - on_cancel
            - on_rating
            - on_support
*/

type ActionType string // Defines the ONDC API call. Any actions other than the enumerated actions are not supported by ONDC Protocol

const (
	SearchAction    ActionType = "search"
	SelectAction    ActionType = "select"
	InitAction      ActionType = "init"
	ConfirmAction   ActionType = "confirm"
	UpdateAction    ActionType = "update"
	StatusAction    ActionType = "status"
	TrackAction     ActionType = "track"
	CancelAction    ActionType = "cancel"
	RatingAction    ActionType = "rating"
	SupportAction   ActionType = "support"
	OnSearchAction  ActionType = "on_search"
	OnSelectAction  ActionType = "on_select"
	OnInitAction    ActionType = "on_init"
	OnConfirmAction ActionType = "on_confirm"
	OnUpdateAction  ActionType = "on_update"
	OnStatusAction  ActionType = "on_status"
	OnTrackAction   ActionType = "on_track"
	OnCancelAction  ActionType = "on_cancel"
	OnRatingAction  ActionType = "on_rating"
	OnSupportAction ActionType = "on_support"
)
