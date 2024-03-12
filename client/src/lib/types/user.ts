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

  constructor(user: UserStyle) {
    this.id = uuidv4()
    this.username = user.username;
    this.email = user.email;
    this.password = user.password;
    this.repeatPassword = user.repeatPassword;
    this.dateOfBirth = user.dateOfBirth;
    this.firstName = user.firstName;
    this.lastName = user.lastName;

    this.initialize()


  }

  private initialize(): void {
    this.firstName = this.capitalizeFirstLetter(this.firstName)
    this.lastName = this.capitalizeFirstLetter(this.lastName)
    const salt = bcrypt.genSaltSync(10);
    console.log(salt)

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
  }
}