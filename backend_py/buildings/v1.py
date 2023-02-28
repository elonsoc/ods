from flask import Blueprint, request
from uuid import uuid4
from enum import Enum

endpoint = Blueprint("buildings", __name__, template_folder="templates", 
    url_prefix="/buildings/v1")

class BuildingType(Enum):
    """defines the different types of a building."""
    DINING_HALL = 0
    CLASSROOM = 1
    OFFICE = 2
    OTHER = 3


class Building:
    """Defines the building class."""
    def __init__(self, name: str, address: str):
        self.name = name
        self.address = address
        self.location = (0, 0)
        self.type = BuildingType.CLASSROOM.value
        self.id = uuid4()
    def to_dict(self):
        return vars(self)

# this represents a databse that stores buildings
buildings_list = dict()
mcewen = Building("McEwen Dining Hall", "Somewhere on Campus")
buildings_list[mcewen.id] = mcewen

@endpoint.route("/", methods=["POST", "GET"])
def index():
    if request.method == "GET":
        
        return buildings_list[mcewen.id].to_dict(), 200
    return "Hello, World!", 200
