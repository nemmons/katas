# Babysitter Kata

## Background
This kata simulates a babysitter working and getting paid for one night.  The rules are pretty straight forward.

## Feature
*As a babysitter<br>
In order to get paid for 1 night of work<br>
I want to calculate my nightly charge<br>*

## Requirements
The babysitter:
- starts no earlier than 5:00PM
- leaves no later than 4:00AM
- only babysits for one family per night
- gets paid for full hours (no fractional hours)
- should be prevented from mistakes when entering times (e.g. end time before start time, or outside of allowable work hours)

The job:
- Pays different rates for each family (based on bedtimes, kids and pets, etc...)
- Family A pays $15 per hour before 11pm, and $20 per hour the rest of the night
- Family B pays $12 per hour before 10pm, $8 between 10 and 12, and $16 the rest of the night
- Family C pays $21 per hour before 9pm, then $15 the rest of the night
- The time ranges are the same as the babysitter (5pm through 4am)

Deliverable:
- Calculate total pay, based on babysitter start and end time, and a family.

## Running
Coming Soon!

## Manual Calculations

| Number | Flat | A   | B   | C   |
|:-------|:-----|:----|:----|:----|
| 5      | 1    | 15  | 12  | 21  |
| 6      | 1    | 15  | 12  | 21  |
| 7      | 1    | 15  | 12  | 21  |
| 8      | 1    | 15  | 12  | 21  |
| 9      | 1    | 15  | 12  | 15  |
| 10     | 1    | 15  | 8   | 15  |
| 11     | 1    | 20  | 8   | 15  |
| 12     | 1    | 20  | 16  | 15  |
| 1      | 1    | 20  | 16  | 15  |
| 2      | 1    | 20  | 16  | 15  |
| 3      | 1    | 20  | 16  | 15  |
| 4      |      |     |     |     |
|        | 11   | 190 | 140 | 189 |

## Attribution

This exercise was lifted from https://github.com/ManApart/babysitter-kata