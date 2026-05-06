import { ref, reactive } from 'vue'

/**
 * Composable for form validation
 * @param {Object} schema - Validation schema with rules
 * @param {Object} initialValues - Initial form values
 */
export function useFormValidation(schema = {}, initialValues = {}) {
  const form = reactive({ ...initialValues })
  const errors = ref({})
  const touched = ref({})
  const isValid = ref(true)

  /**
   * Validate a single field
   */
  const validateField = (fieldName, value) => {
    const rules = schema[fieldName]
    if (!rules) return true

    const fieldErrors = []

    // Required validation
    if (rules.required && !value) {
      fieldErrors.push(rules.requiredMessage || 'This field is required')
    }

    // Min length validation
    if (rules.minLength && value && value.length < rules.minLength) {
      fieldErrors.push(rules.minLengthMessage || `Minimum ${rules.minLength} characters required`)
    }

    // Max length validation
    if (rules.maxLength && value && value.length > rules.maxLength) {
      fieldErrors.push(rules.maxLengthMessage || `Maximum ${rules.maxLength} characters allowed`)
    }

    // Email validation
    if (rules.email && value) {
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      if (!emailRegex.test(value)) {
        fieldErrors.push(rules.emailMessage || 'Please enter a valid email')
      }
    }

    // Pattern validation
    if (rules.pattern && value && !rules.pattern.test(value)) {
      fieldErrors.push(rules.patternMessage || 'Invalid format')
    }

    // Custom validation
    if (rules.custom && typeof rules.custom === 'function') {
      const customError = rules.custom(value, form)
      if (customError) {
        fieldErrors.push(customError)
      }
    }

    // Set errors for this field
    if (fieldErrors.length > 0) {
      errors.value[fieldName] = fieldErrors[0]
      return false
    } else {
      delete errors.value[fieldName]
      return true
    }
  }

  /**
   * Validate all fields
   */
  const validate = () => {
    let allValid = true
    Object.keys(schema).forEach(fieldName => {
      const fieldValid = validateField(fieldName, form[fieldName])
      if (!fieldValid) allValid = false
    })
    isValid.value = allValid
    return allValid
  }

  /**
   * Validate field on input
   */
  const handleInput = (fieldName) => {
    validateField(fieldName, form[fieldName])
  }

  /**
   * Mark field as touched on blur
   */
  const handleBlur = (fieldName) => {
    touched.value[fieldName] = true
    validateField(fieldName, form[fieldName])
  }

  /**
   * Reset form
   */
  const resetForm = () => {
    Object.keys(initialValues).forEach(key => {
      form[key] = initialValues[key]
    })
    errors.value = {}
    touched.value = {}
    isValid.value = true
  }

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
