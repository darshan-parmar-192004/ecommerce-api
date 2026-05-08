export const useFormValidation = (schema = {}, initialValues = {}) => {
  const form = ref({ ...initialValues })
  const errors = ref({})
  const touched = ref({})

  const validate = () => {
    const newErrors = {}
    for (const [field, rules] of Object.entries(schema)) {
      const value = form.value[field]
      if (rules.required && (!value || (typeof value === 'string' && !value.trim()))) {
        newErrors[field] = `${field} is required`
      } else if (rules.minLength && value && value.length < rules.minLength) {
        newErrors[field] = `${field} must be at least ${rules.minLength} characters`
      } else if (rules.maxLength && value && value.length > rules.maxLength) {
        newErrors[field] = `${field} must be at most ${rules.maxLength} characters`
      } else if (rules.email && value && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) {
        newErrors[field] = 'Invalid email address'
      } else if (rules.pattern && value && !rules.pattern.test(value)) {
        newErrors[field] = rules.patternMessage || 'Invalid format'
      } else if (rules.custom && typeof rules.custom === 'function' && !rules.custom(value, form.value)) {
        newErrors[field] = rules.customMessage || 'Invalid value'
      }
    }
    errors.value = newErrors
    return Object.keys(newErrors).length === 0
  }

  const validateField = (field, value) => {
    form.value[field] = value
    const rules = schema[field]
    if (!rules) return true

    let error = null
    if (rules.required && (!value || (typeof value === 'string' && !value.trim()))) {
      error = `${field} is required`
    } else if (rules.minLength && value && value.length < rules.minLength) {
      error = `${field} must be at least ${rules.minLength} characters`
    } else if (rules.maxLength && value && value.length > rules.maxLength) {
      error = `${field} must be at most ${rules.maxLength} characters`
    } else if (rules.email && value && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) {
      error = 'Invalid email address'
    } else if (rules.pattern && value && !rules.pattern.test(value)) {
      error = rules.patternMessage || 'Invalid format'
    } else if (rules.custom && typeof rules.custom === 'function' && !rules.custom(value, form.value)) {
      error = rules.customMessage || 'Invalid value'
    }

    if (error) {
      errors.value = { ...errors.value, [field]: error }
    } else {
      const { [field]: _, ...rest } = errors.value
      errors.value = rest
    }
    return !error
  }

  const handleInput = (field, value) => {
    form.value[field] = value
    if (touched.value[field]) {
      validateField(field, value)
    }
  }

  const handleBlur = (field) => {
    touched.value[field] = true
    validateField(field, form.value[field])
  }

  const resetForm = () => {
    form.value = { ...initialValues }
    errors.value = {}
    touched.value = {}
  }

  const isValid = computed(() => {
    return Object.keys(errors.value).length === 0
  })

  return {
    form,
    errors,
    touched,
    isValid,
    validate,
    validateField,
    handleInput,
    handleBlur,
    resetForm
  }
}