import { ref, onMounted } from 'vue'

const MOCK_PRODUCTS = [
  {
    id: 1,
    name: 'Wireless Headphones',
    price: 79.99,
    category: 'Electronics',
    description: 'Premium noise-canceling wireless headphones with 30-hour battery life and crystal clear audio quality.'
  },
  {
    id: 2,
    name: 'Ergonomic Office Chair',
    price: 299.99,
    category: 'Furniture',
    description: 'Adjustable lumbar support, breathable mesh back, and premium foam cushioning for all-day comfort.'
  },
  {
    id: 3,
    name: 'Smart Watch Pro',
    price: 249.99,
    category: 'Electronics',
    description: 'Advanced health tracking, GPS, water resistance up to 50m, and seamless notifications integration.'
  },
  {
    id: 4,
    name: 'Minimalist Desk Lamp',
    price: 45.99,
    category: 'Lighting',
    description: 'Touch-controlled LED desk lamp with adjustable brightness and color temperature for focused work.'
  },
  {
    id: 5,
    name: 'Mechanical Keyboard',
    price: 129.99,
    category: 'Electronics',
    description: 'RGB backlit mechanical keyboard with Cherry MX switches, programmable macros, and aluminum frame.'
  },
  {
    id: 6,
    name: 'Standing Desk Converter',
    price: 189.99,
    category: 'Furniture',
    description: 'Height-adjustable workstation that transforms any desk into a sit-stand desk with gas spring lift.'
  },
  {
    id: 7,
    name: 'Ultra-Wide Monitor',
    price: 449.99,
    category: 'Electronics',
    description: '34-inch curved ultrawide monitor with 144Hz refresh rate and HDR support for immersive productivity.'
  },
  {
    id: 8,
    name: 'Wireless Mouse',
    price: 49.99,
    category: 'Electronics',
    description: 'Ergonomic wireless mouse with precision tracking and customizable buttons for maximum efficiency.'
  }
]

export function useProducts(apiUrl = 'http://localhost:3000/api') {
  const products = ref([])
  const loading = ref(false)
  const error = ref(null)

  const fetchProducts = async () => {
    loading.value = true
    error.value = null
    
    try {
      const response = await fetch(`${apiUrl}/products`)
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      
      const data = await response.json()
      products.value = data
    } catch (e) {
      console.warn('API unavailable, using mock data:', e.message)
      products.value = MOCK_PRODUCTS
    } finally {
      loading.value = false
    }
  }

  const refetch = () => {
    fetchProducts()
  }

  onMounted(() => {
    fetchProducts()
  })

  return {
    products,
    loading,
    error,
    refetch
  }
}
