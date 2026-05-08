const snakeToCamel = (str) =>
  str.replace(/_([a-z])/g, (_, letter) => letter.toUpperCase())

const camelToSnake = (str) =>
  str.replace(/[A-Z]/g, (letter) => `_${letter.toLowerCase()}`)

const isObject = (val) =>
  val !== null && typeof val === 'object' && !Array.isArray(val)

export const mapKeys = (obj, transform) => {
  if (!obj || typeof obj !== 'object') return obj
  if (Array.isArray(obj)) return obj.map((item) => mapKeys(item, transform))
  return Object.keys(obj).reduce((acc, key) => {
    const transformedKey = transform(key)
    acc[transformedKey] = mapKeys(obj[key], transform)
    return acc
  }, {})
}

export const snakeToCamelCase = (obj) => mapKeys(obj, snakeToCamel)

export const camelToSnakeCase = (obj) => mapKeys(obj, camelToSnake)
