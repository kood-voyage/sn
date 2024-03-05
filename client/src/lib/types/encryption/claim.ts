
export class Claims {
  claimsMap: Record<string, string | number>

  constructor() {
    this.claimsMap = {}

    const time = new Date()
    this.setTime("iat", time)
  }



  set(key: string, val: string) {
    this.claimsMap[key] = val
  }

  setTime(key: string, val: Date) {
    this.claimsMap[key] =  Math.floor(val.getTime() / 1000)
  }

  get(key: string) {
    const val = this.claimsMap[key]
    if (typeof val === "string") return val
    if (typeof val === "number") return val
    return val
  }

  getTime(key: string) {
    const val = this.claimsMap[key]
    if (val == undefined) return {ok: false, msg: "No such key in the claims map!"}

    return new Date(val)
  }

  hasClaim(key: string) {
    if (this.claimsMap[key] !== undefined) return true
    return false
  }

}

