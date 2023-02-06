import { ref } from "vue";
import emailApi from "@/api/emailApi";

const useEmails = () => {
  const emails = ref([]);

  const getEmails = async () => {
    const { data } = await emailApi.get(`/emails`);

    emails.value = data.hits.hits?.map(
      (email: { _id: string; _source: object }) => {
        return { id: email._id, source: email._source };
      }
    );
  };

  const searchTerm = async (term: string) => {
    const body = {
      Term: term,
    };
    const { data } = await emailApi.post(`/emails/term`, body);
    emails.value = data.hits.hits?.map(
      (email: { _id: string; _source: object }) => {
        return { id: email._id, source: email._source };
      }
    );
  };

  getEmails();

  return { emails, getEmails, searchTerm };
};

export default useEmails;
