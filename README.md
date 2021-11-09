# Prioritized

## Features
- [x] Categories / Adjustable Times / Sub Types
- [x] Task Manager / Category Groupings
- [x] Auto Scheduler / Based on Task and Category importance

# Prioritized API Documentation
## Endpoints
> 
* **/insert**: Used to insert tasks into a task grouping
* **/sort**: Sort a task Grouping

## Data Structures
---
### Task Grouping: 
```
type: json
```
```
{
	"name": "<INSERT GROUPING NAME>",
	"time_ranges": {  # (REQUIRED)
		"time_start": "0000-01-01Thh:mm:ss+07:00(TIMEZONE OFFSET)",
		"time_end": "0000-01-01Thh:mm:ss+07:00(TIMEZONE OFFSET)"
	}
	"weekdays": [1,2,3,4,5] (Days of week, 0 = Sunday & 6 = Sturday, empty list [] for everyday),
	"weight_coef": 1.2 # Any floating point number between 1.0 and 2.0 (REQUIRED)
	"tasks": [
		<INSERT TASK OBJECT>,
		<INSERT TASK OBJECT>,
		<INSERT TASK OBJECT>
	]
	
}
```
---
### Task:
	
```
{
	"name": "<INSERT TASK NAME>", # (REQUIRED)
	"timeline": {
		"time_start": "yyyy-mm-ddThh:mm:ss+07:00(TIMEZONE OFFSET)",
		"time_end": "yyyy-mm-ddThh:mm:ss+07:00(TIMEZONE OFFSET)"
	},
	"assigned_time": {
		"time_start": "yyyy-mm-ddThh:mm:ss+07:00(TIMEZONE OFFSET)",
		"time_end": "yyyy-mm-ddThh:mm:ss+07:00(TIMEZONE OFFSET)"
	},
	"fixed": false,
	"weight_coef": 1.2 # Any floating point number between 1.0 and 2.0 # (REQUIRED)
	"estimated_time": "1h30m0s", # (REQUIRED)
	"current_score": 2000
},
```

## Requests body examples:

### To /insert:
```
{
    "time_preference": 30, # User time preference for time on task (by default pass 30)
    "task_grouping": {
        "weight_coef": 1.0, # Grouping weight coeffiecient (must be between 1 and 2)
        "time_ranges": [ # Time range to assign tasks
            {
                "time_start": "0000-01-01T09:00:00+07:00",
                "time_end": "0000-01-01T12:00:00+07:00"
            },
            {
                "time_start": "0000-01-01T13:00:00+07:00",
                "time_end": "0000-01-01T17:00:00+07:00"
            }
        ],
        "weekdays": [1,2,3,4,5], # Monday through friday
        "tasks": [ # Task objects that have been previously assigned
            {
                "name": "Do Task1",
                "timeline": {
                    "time_start": "2021-10-01T00:00:00+07:00",
                    "time_end": "2021-11-10T00:00:00+07:00"
                },
                "assigned_time": {
                    "time_start": "2021-11-21T12:00:00+07:00",
                    "time_end": "2021-11-21T12:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 2000
            },
            {
                "name": "Do Task2",
                "assigned_time": {
                    "time_start": "2021-09-21T12:00:00+07:00",
                    "time_end": "2021-11-21T12:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 1800
            },
            {
                "name": "Do Task3",
                "assigned_time": {
                    "time_start": "2021-09-21T12:00:00+07:00",
                    "time_end": "2021-09-21T14:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 2000
            },
            {
                "name": "Do Task5",
                "assigned_time": {
                    "time_start": "2021-09-21T12:00:00+07:00",
                    "time_end": "2021-09-21T14:00:00+07:00"
                },
                "timeline": {
                    "time_end": "2021-05-01T00:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 2523
            },
            {
                "name": "Do Task6",
                "assigned_time": {
                    "time_start": "2021-09-21T12:00:00+07:00",
                    "time_end": "2021-09-21T14:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 2069
            }
        ]
    },
    "insert_tasks": [ # New task objects to be added into "tasks"
        {
            "name": "Do Task4",
            "timeline": { # Timeline with only deadline
                "time_end": "2021-05-01T00:00:00+07:00"
            },
            "weight_coef": 1.7, # Task weight coefficient
            "estimated_time": "90m" # Estimated time on task "10h10m10s"
        },
        {
             "name": "Do Task7",
            "timeline": {
                "time_end": "2021-11-30T00:00:00+07:00"
            },
            "weight_coef": 1.5,
            "estimated_time": "60m"
        }
    ]
}
```

### To /sort:
```
{
    "task_grouping": { # Send one task grouping
        "weight_coef": 1.0, 
        "time_ranges": [ # Time range to assign tasks
            {
                "time_start": "0000-01-01T09:00:00+07:00",
                "time_end": "0000-01-01T12:00:00+07:00"
            },
            {
                "time_start": "0000-01-01T13:00:00+07:00",
                "time_end": "0000-01-01T17:00:00+07:00"
            }
        ],
        "weekdays": [1,2,3,4,5], # Monday through Friday
        "tasks": [ # Tasks of the grouping
            {
                "name": "Do Task1",
                "timeline": {
                    "time_start": "2021-10-01T00:00:00+07:00",
                    "time_end": "2021-11-10T00:00:00+07:00"
                },
                "assigned_time": {
                    "time_start": "2021-11-21T12:00:00+07:00",
                    "time_end": "2021-11-21T12:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 2000
            },
            {
                "name": "Do Task2",
                "assigned_time": {
                    "time_start": "2021-09-21T12:00:00+07:00",
                    "time_end": "2021-11-21T12:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 1800
            },
            {
                "name": "Do Task3",
                "assigned_time": {
                    "time_start": "2021-09-21T12:00:00+07:00",
                    "time_end": "2021-09-21T14:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 2000
            },
            {
                "name": "Do Task5",
                "assigned_time": {
                    "time_start": "2021-09-21T12:00:00+07:00",
                    "time_end": "2021-09-21T14:00:00+07:00"
                },
                "timeline": {
                    "time_end": "2021-05-01T00:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 2523
            },
            {
                "name": "Do Task6",
                "assigned_time": {
                    "time_start": "2021-09-21T12:00:00+07:00",
                    "time_end": "2021-09-21T14:00:00+07:00"
                },
                "fixed": false,
                "estimated_time": "1h30m0s",
                "current_score": 2069
            }
        ]
    }
}
```

### Response
**Both endpoints with have the same response (A list of tasks that has been scored, sorted and assigned with a time)**

```
[
    {
        "name": "Do Task1",
        "timeline": {
            "time_start": "2021-10-01T00:00:00+07:00",
            "time_end": "2021-11-10T00:00:00+07:00"
        },
        "assigned_time": {
            "time_start": "2021-11-10T09:00:00+07:00",
            "time_end": "2021-11-10T10:30:00+07:00"
        },
        "estimated_time": "1h30m0s",
        "current_score": 2000
    },
    {
        "name": "Do Task4",
        "timeline": {
            "time_start": "0001-01-01T00:00:00Z",
            "time_end": "2021-05-01T00:00:00+07:00"
        },
        "assigned_time": {
            "time_start": "2021-11-10T10:30:00+07:00",
            "time_end": "2021-11-10T11:30:00+07:00"
        },
        "estimated_time": "1h30m0s",
        "weight_coef": 1.7,
        "current_score": 2125
    },
    {
        "name": "Do Task7",
        "timeline": {
            "time_start": "0001-01-01T00:00:00Z",
            "time_end": "2021-11-30T00:00:00+07:00"
        },
        "assigned_time": {
            "time_start": "2021-11-10T13:00:00+07:00",
            "time_end": "2021-11-10T14:30:00+07:00"
        },
        "estimated_time": "1h0m0s",
        "weight_coef": 1.5,
        "current_score": 1125
    },
    {
        "name": "Do Task6",
        "timeline": {
            "time_start": "0001-01-01T00:00:00Z",
            "time_end": "0001-01-01T00:00:00Z"
        },
        "assigned_time": {
            "time_start": "2021-11-10T14:30:00+07:00",
            "time_end": "2021-11-10T16:00:00+07:00"
        },
        "estimated_time": "1h30m0s",
        "current_score": 2069
    },
    {
        "name": "Do Task2",
        "timeline": {
            "time_start": "0001-01-01T00:00:00Z",
            "time_end": "0001-01-01T00:00:00Z"
        },
        "assigned_time": {
            "time_start": "2021-11-11T09:00:00+07:00",
            "time_end": "2021-11-11T10:30:00+07:00"
        },
        "estimated_time": "1h30m0s",
        "current_score": 1800
    },
    {
        "name": "Do Task3",
        "timeline": {
            "time_start": "0001-01-01T00:00:00Z",
            "time_end": "0001-01-01T00:00:00Z"
        },
        "assigned_time": {
            "time_start": "2021-11-11T13:00:00+07:00",
            "time_end": "2021-11-11T14:30:00+07:00"
        },
        "estimated_time": "1h30m0s",
        "current_score": 2000
    },
    {
        "name": "Do Task5",
        "timeline": {
            "time_start": "0001-01-01T00:00:00Z",
            "time_end": "2021-05-01T00:00:00+07:00"
        },
        "assigned_time": {
            "time_start": "2021-11-11T14:30:00+07:00",
            "time_end": "2021-11-11T16:00:00+07:00"
        },
        "estimated_time": "1h30m0s",
        "current_score": 2523
    }
]
```

## Notes:
1. Note that time_ranges have yyyy-mm-dd are set as 0000-01-01 maintain this for all representations above. Only fill yyyy-mm-dd if the example specifies yyyy-mm-dd.
2. Make sure that the timezone offset is set correctly, or else there may be unpredictable behavior.
3. In the Data Structure section, watch for comments with "(REQUIRED)". Any requests to the server missing these fields will result in an error and the request will not be processed. All other fields can be missing.
