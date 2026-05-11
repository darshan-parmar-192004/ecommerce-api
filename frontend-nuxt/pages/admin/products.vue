<script setup>
import { useApi } from '~/composables/useApi'

definePageMeta({
  layout: 'admin',
  middleware: 'admin'
})

const products = ref([])
const loading = ref(true)
const showModal = ref(false)
const editingProduct = ref(null)

const form = ref({ name: '', price: 0, category_id: '', description: '', stock: 0 })
const { admin: adminApi } = useApi()

const fetchProducts = async () => {
  loading.value = true
  try {
    const response = await adminApi.products.list()
    products.value = response.data || response || []
  } catch (err) {
    console.error('Failed to fetch products', err)
  } finally {
    loading.value = false
  }
}

const editProduct = (product) => {
  editingProduct.value = product
  form.value = { ...product }
  showModal.value = true
}

const saveProduct = async () => {
  try {
    if (editingProduct.value) {
      await adminApi.products.update(editingProduct.value.product_id, form.value)
    } else {
      await adminApi.products.create(form.value)
    }
    showModal.value = false
    await fetchProducts()
  } catch (err) {
    console.error('Failed to save product', err)
  }
}

const resetForm = () => {
  editingProduct.value = null
  form.value = { name: '', price: 0, category_id: '', description: '', stock: 0 }
}

onMounted(fetchProducts)
</script>

<template>
  <div class="p-6">
    <div class="flex items-center justify-between mb-8">
      <h1 class="text-3xl font-bold text-on_surface">Manage Products</h1>
      <Button @click="showModal = true; resetForm()" label="Add Product" icon="pi pi-plus" />
    </div>

    <div v-if="loading" class="space-y-3">
      <div v-for="n in 5" :key="n" class="h-16 bg-surface-container rounded-lg animate-pulse" />
    </div>

    <div v-else class="bg-surface-container-lowest rounded-xl border border-outline-variant/20 overflow-hidden">
      <DataTable :value="products" stripedRows class="w-full">
        <Column field="name" header="Name" />
        <Column header="Price">
          <template #body="{ data }">${{ data.price?.toFixed(2) }}</template>
        </Column>
        <Column field="stock" header="Stock" />
        <Column header="Actions">
          <template #body="{ data }">
            <Button @click="editProduct(data)" icon="pi pi-pencil" text rounded size="small" />
          </template>
        </Column>
      </DataTable>
    </div>

    <Dialog v-model:visible="showModal" :header="editingProduct ? 'Edit Product' : 'Add Product'" modal class="w-full max-w-md">
      <div class="space-y-4 p-4">
        <div>
          <label class="block text-sm font-medium text-on_surface mb-1">Name</label>
          <InputText v-model="form.name" class="w-full" placeholder="Product Name" />
        </div>
        <div>
          <label class="block text-sm font-medium text-on_surface mb-1">Price</label>
          <InputNumber v-model="form.price" mode="currency" currency="USD" locale="en-US" class="w-full" />
        </div>
        <div>
          <label class="block text-sm font-medium text-on_surface mb-1">Stock</label>
          <InputNumber v-model="form.stock" class="w-full" />
        </div>
      </div>
      <template #footer>
        <Button @click="showModal = false" label="Cancel" text />
        <Button @click="saveProduct" label="Save" icon="pi pi-check" />
      </template>
    </Dialog>
  </div>
</template>
