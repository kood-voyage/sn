import bcrypt from 'bcrypt'
import { v4 as uuidv4 } from 'uuid';

export type UserStyle = {
  username: string,
  email: string,
  password: string,
  repeatPassword: string,
  dateOfBirth: string,
  firstName: string,
  lastName: string,
  avatar: string,
  cover: string,
  description: string
}

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

export class User implements UserStyle {
  id: string;
  username: string;
  email: string;
  password: string;
  repeatPassword: string;
  dateOfBirth: string;
  firstName: string;
  lastName: string;
  avatar: string;
  cover: string;
  description: string;


  constructor(user: UserStyle) {
    this.id = uuidv4()
    this.username = user.username;
    this.email = user.email;
    this.password = user.password;
    this.repeatPassword = user.repeatPassword;
    this.dateOfBirth = user.dateOfBirth;
    this.firstName = user.firstName;
    this.lastName = user.lastName;
    this.avatar = user.avatar;
    this.cover = user.cover;
    this.description = user.description;

    this.initialize()
  }

  private initialize(): void {
    this.firstName = this.capitalizeFirstLetter(this.firstName)
    this.lastName = this.capitalizeFirstLetter(this.lastName)

    this.avatar = `https://api.dicebear.com/7.x/bottts-neutral/svg?seed=` + this.username
    this.cover = "https://media.wired.com/photos/61f48f02d0e55ccbebd52d15/3:2/w_2400,h_1600,c_limit/Gear-Rant-Game-Family-Plans-1334436001.jpg"
    this.description = ""

    const salt = bcrypt.genSaltSync(10);

    const hash = bcrypt.hashSync(this.password, salt);
    this.password = hash
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

  }
}