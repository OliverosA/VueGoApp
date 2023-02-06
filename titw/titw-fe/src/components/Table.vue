<template>
  <div class="justify-start mt-10 w-1/2 mx-3">
    <div
      v-if="!emails || emails.length == 0"
      class="bg-blue-100 rounded-lg py-5 px-6 mb-3 text-base text-blue-700 inline-flex items-center w-full"
      role="alert"
    >
      <svg
        aria-hidden="true"
        focusable="false"
        data-prefix="fas"
        data-icon="info-circle"
        class="w-4 h-4 mr-2 fill-current"
        role="img"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 512 512"
      >
        <path
          fill="currentColor"
          d="M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 110c23.196 0 42 18.804 42 42s-18.804 42-42 42-42-18.804-42-42 18.804-42 42-42zm56 254c0 6.627-5.373 12-12 12h-88c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h12v-64h-12c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h64c6.627 0 12 5.373 12 12v100h12c6.627 0 12 5.373 12 12v24z"
        ></path>
      </svg>
      No Emails Available
    </div>
    <div v-else class="tableContainer">
      <table class="table-fixed border-collapse w-full max-h-44 text-center">
        <thead class="border-b bg-gray-800">
          <tr>
            <th
              class="p-3 font-bold uppercase bg-gray-200 text-gray-600 border border-gray-300 px-6 py-4 hidden lg:table-cell"
            >
              Subject
            </th>
            <th
              class="p-3 font-bold uppercase bg-gray-200 text-gray-600 border border-gray-300 px-6 py-4 hidden lg:table-cell"
            >
              From
            </th>
            <th
              class="p-3 font-bold uppercase bg-gray-200 text-gray-600 border border-gray-300 px-6 py-4 hidden lg:table-cell"
            >
              To
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            class="bg-white lg:hover:bg-sky-200 h-10"
            v-for="{ id, source } in emails"
            :key="id"
            @click="
              $emit('selectedEmail', {
                id: id,
                subject: source.subject,
                from: source.from,
                to: source.to,
                content: source.content,
              })
            "
          >
            <td class="truncate text-gray-800 text-center border border-b">
              {{ source.subject }}
            </td>
            <td class="truncate text-gray-800 text-center border border-b">
              {{ source.from }}
            </td>
            <td class="truncate text-gray-800 text-center border border-b">
              {{ source.to }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  props: {
    emails: Array,
  },
};
</script>

<style scoped>
.tableContainer {
  height: 80vh;
  overflow-y: scroll;
  text-align: center;
}
</style>
