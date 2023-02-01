import axios from "axios";
import { ref } from "vue";

const useEmails = () => {
  const emails = ref([]);

  const getEmails = async () => {
    const { data } = await axios.get(`http://localhost:3000/emails`);

    emails.value = data.hits.hits?.map((email: { _id: any; _source: any }) => {
      return { id: email._id, source: email._source };
    });
  };

  const searchTerm = async (term: string) => {
    const body = {
      Term: term,
    };
    const { data } = await axios.post(
      `http://localhost:3000/emails/term`,
      body
    );
    emails.value = data.hits.hits?.map((email: { _id: any; _source: any }) => {
      return { id: email._id, source: email._source };
    });
  };

  getEmails();

  return { emails, getEmails, searchTerm };
};

export default useEmails;
