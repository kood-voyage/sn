
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
  username: string;
  email: string;
  password: string;
  repeatPassword: string;
  dateOfBirth: string;
  firstName: string;
  lastName: string;

  constructor(user: UserStyle) {
    this.username = user.username;
    this.email = user.email;
    this.password = user.password;
    this.repeatPassword = user.repeatPassword;
    this.dateOfBirth = user.dateOfBirth;
    this.firstName = user.firstName;
    this.lastName = user.lastName;
  }

  *[Symbol.iterator]() {
    yield this.username;
    yield this.email;
    yield this.password;
    yield this.dateOfBirth;
    yield this.firstName;
    yield this.lastName;
  }
}