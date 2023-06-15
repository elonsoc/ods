export const HOME_PAGE_COURSE_SAMPLE = `{
    course_id: 1234,
    name: Computer Science II,
    description: This course continues the study of object-oriented programming with an emphasis on graphical user interfaces, event handling...,
    instructor: Dr. Professor,
    schedule: [
      {
        day: Monday,
        start_time: 12:30 PM,
        end_time: 1:40 PM
      },
      {
        day: Wednesday,
        start_time: 12:30 PM,
        end_time: 1:40 PM
      },
    ],
    rotation : {
        semesters: [Fall, Spring]
        periodical: yearly
    },
    location: Mooney Building, Room 202,
    credits: 4,
    prerequisites: [CSC 1300]
}`;

export const QUERY_PARAMETER_SECTION = `const location = 'Main Campus';
const capacity = 100;

const queryString = \`location= \${encodeURIComponent(location)}&capacity= \${encodeURIComponent(capacity)}\`;
const url = \`http://api.ods.elon.edu/v1/buildings? \${queryString}\`;

// Make the GET request to the URL
`;

export const REQUEST_PAYLOAD_SAMPLE = `{
  "id": "08ac60ab-e415-4d2c-b274-869e90a8725e"
}`;

export const RESPONSE_PAYLOAD_SAMPLE = `{
  "id": "08ac60ab-e415-4d2c-b274-869e90a8725e",
  "name": "Smith Hall",
  "location": "Main Campus",
  "capacity": 100,
  "floors": 3
}
`;

export const BUILDING_RESPONSE_SAMPLE = `{"mcewen": {
  Name: "McEwen Dining Hall",
  Floors: []Floor{
    {Name: "Floor 1", Level: 1, Rooms: []Room{ { Name: "Room 1", Level: 1}, { Name: "Room 2", Level: 1}}},
    {Name: "Floor 2", Level: 2, Rooms: []Room{ { Name: "Room 3", Level: 2}, { Name: "Room 4", Level: 2}}},
  },
  Location:     LatLng{ Lat: 37.422, Lng: -122.084 },
  Address:      "1600 Amphitheatre Parkway, Mountain View, CA 94043",
  BuildingType: BuildingTypeDining,
  Id:           "mcewen",
}}`;
