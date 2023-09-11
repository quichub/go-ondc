package schema

/*
   Time:
     description: Describes time in its various forms. It can be a single point in time; duration; or a structured timetable of operations
     type: object
     properties:
       label:
         type: string
       timestamp:
         type: string
         format: date-time
       duration:
         $ref: '#/components/schemas/Duration'
       range:
         type: object
         properties:
           start:
             type: string
             format: date-time
           end:
             type: string
             format: date-time
       days:
         type: string
         description: comma separated values representing days of the week
       schedule:
         $ref: '#/components/schemas/Schedule'
*/

type TimeRange struct {
	Start DateTime `json:"start,omitempty"` // Timestamp in RFC3339 format
	End   DateTime `json:"end,omitempty"`   // Timestamp in RFC3339 format
}

type Time struct {
	Label     string    `json:"label,omitempty"`     // Describes time in its various forms. It can be a single point in time; duration; or a structured timetable of operations
	Timestamp DateTime  `json:"timestamp,omitempty"` // Timestamp in RFC3339 format
	Duration  Duration  `json:"duration,omitempty"`  // Describes a duration
	Range     TimeRange `json:"range,omitempty"`     // Describes a range
	Days      string    `json:"days,omitempty"`      // comma separated values representing days of the week
	Schedule  Schedule  `json:"schedule,omitempty"`  // Describes a schedule
}
