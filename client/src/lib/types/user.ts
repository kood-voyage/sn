
import { v4 as uuidv4 } from 'uuid';





export type UserModel = {
  id: string,
  username: string,
  email: string,
  date_of_birth: string,
  first_name: string,
  last_name: string,
  avatar: string,
  cover: string,
  description: string
}


// 	{
//     "id": "somerandomidlaterfromfrontend",
//     "username": "testuser",
//     "email": "test@user.ee",
//     "password": "testUser",
//     "timestamp": "2024-04-25T00:00:00Z",
//     "date_of_birth": "11-11-1111",
//     "first_name": "first name",
//     "last_name": "last name",
//     "description": "this is my fantastic bio",
//     "avatar": "your_avatar_url_here",
//     "cover": "your_cover_url_here",
//     "member_type": "public"
// }

export type UserType = {
  id: string;
  username: string,
  email: string,
  password: string,
  dateOfBirth: string,
  firstName: string,
  lastName: string,
  avatar?: string,
  cover?: string,
  description?: string,
  event_status: string
}

export class CreateUser implements UserType {
  id: string;
  username: string;
  email: string;
  password: string;
  dateOfBirth: string;
  firstName: string;
  lastName: string;
  avatar?: string;
  cover?: string;
  description?: string;
  event_status: string;


  constructor(user: UserType) {
    this.id = uuidv4()
    this.username = user.username;
    this.email = user.email;
    this.password = user.password;
    this.dateOfBirth = user.dateOfBirth;
    this.firstName = user.firstName;
    this.lastName = user.lastName;
    this.avatar = user.avatar;
    this.cover = user.cover;
    this.description = user.description;
    this.event_status = user.event_status

    this.initialize()
  }

  private initialize(): void {
    this.firstName = this.capitalizeFirstLetter(this.firstName)
    this.lastName = this.capitalizeFirstLetter(this.lastName)

    this.avatar = `https://api.dicebear.com/7.x/bottts-neutral/svg?seed=` + this.username
    this.cover = "https://media.wired.com/photos/61f48f02d0e55ccbebd52d15/3:2/w_2400,h_1600,c_limit/Gear-Rant-Game-Family-Plans-1334436001.jpg"
    this.description = ""


  }

  private capitalizeFirstLetter(string: string) {
    return string.charAt(0).toUpperCase() + string.slice(1).toLowerCase();
  }


  *[Symbol.iterator]() {
    yield this.id;
    yield this.username;
    yield this.email;
    yield this.password;
    yield this.dateOfBirth;
    yield this.firstName;
    yield this.lastName;
    yield this.avatar;
    yield this.cover;
    yield this.description
    yield this.event_status

  }
}
