<template>
  <Input @get-input-value="getInputValue" />
  <div class="flex justify-center w-full">
    <Table :emails="emails" @selected-email="getSelectedEmail" />
    <ContentEmail :selectedEmail="selectedEmail" />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import Table from "@/components/Table.vue";
import ContentEmail from "@/components/ContentEmail.vue";
import Input from "@/components/Input.vue";
import useEmails from "@/composables/useEmails";

export default defineComponent({
  name: "HomeView",
  components: { Table, ContentEmail, Input },

  setup() {
    const { emails, searchTerm, getEmails } = useEmails();
    const selectedEmail = reactive({
      id: "",
      subject: "",
      from: "",
      to: "",
      content: "",
    });

    const getInputValue = (inputValue: string) => {
      if (inputValue != "") {
        searchTerm(inputValue);
        return;
      }
      getEmails();
      return;
    };

    const getSelectedEmail = (email: {
      id: string;
      subject: string;
      from: string;
      to: string;
      content: string;
    }) => {
      (selectedEmail.id = email.id),
        (selectedEmail.subject = email.subject),
        (selectedEmail.from = email.from),
        (selectedEmail.to = email.to),
        (selectedEmail.content = email.content);
    };

    return {
      emails,
      getInputValue,
      getSelectedEmail,
      selectedEmail,
    };
  },
});
</script>
