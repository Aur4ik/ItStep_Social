// src/pages/ChatDetailPage.jsx
import { useEffect, useRef, useState } from "react";
import { useParams } from "react-router-dom";
import { getChatMessages, sendMessage, addChatMember } from "../api/chats";
import { useAuth } from "../context/AuthContext";

export default function ChatDetailPage() {
  const { id } = useParams();
  const { user } = useAuth();
  const [messages, setMessages] = useState([]);
  const [text, setText] = useState("");
  const bottomRef = useRef(null);

  useEffect(() => {
    // Вступаем в чат при открытии
    addChatMember(id, { user_id: user?.id }).catch(() => {});

    const load = () => {
      getChatMessages(id)
        .then((res) => setMessages(res.data || []))
        .catch(() => {});
    };

    load();
    const interval = setInterval(load, 3000);
    return () => clearInterval(interval);
  }, [id]);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  const handleSend = async (e) => {
    e.preventDefault();
    if (!text.trim()) return;
    try {
      await sendMessage(id, { content: text });
      setText("");
    } catch {
      alert("Ошибка отправки");
    }
  };

  return (
    <div style={styles.wrapper}>
      <div style={styles.header}>
        <h2>💬 Чат #{id}</h2>
      </div>

      <div style={styles.messages}>
        {messages.length === 0 && (
          <div style={styles.empty}>Сообщений пока нет</div>
        )}
        {messages.map((msg, i) => {
          const isMe = Number(msg.sender_id) === Number(user?.id);
          return (
            <div key={i} style={{ display: "flex", justifyContent: isMe ? "flex-end" : "flex-start", marginBottom: 8 }}>
              <div style={isMe ? styles.myMsg : styles.theirMsg}>
                {!isMe && (
                  <div style={styles.sender}>
                    {msg.sender?.first_name || "Аноним"}
                  </div>
                )}
                <div>{msg.content}</div>
                <div style={styles.time}>
                  {new Date(msg.created_at).toLocaleTimeString("ru", { hour: "2-digit", minute: "2-digit" })}
                </div>
              </div>
            </div>
          );
        })}
        <div ref={bottomRef} />
      </div>

      <form onSubmit={handleSend} style={styles.form}>
        <input
          style={styles.input}
          value={text}
          onChange={(e) => setText(e.target.value)}
          placeholder="Написать сообщение..."
        />
        <button type="submit" style={styles.btn}>➤</button>
      </form>
    </div>
  );
}

const styles = {
  wrapper: { maxWidth: 640, margin: "24px auto", padding: "0 16px", display: "flex", flexDirection: "column", height: "calc(100vh - 80px)" },
  header: { marginBottom: 12 },
  messages: { flex: 1, overflowY: "auto", background: "#fff", borderRadius: 10, padding: 16, boxShadow: "0 2px 8px rgba(0,0,0,0.08)", marginBottom: 12 },
  empty: { textAlign: "center", color: "#999", marginTop: 40 },
  myMsg: { background: "#e94560", color: "#fff", padding: "8px 14px", borderRadius: "16px 16px 4px 16px", maxWidth: 280 },
  theirMsg: { background: "#f4f6fb", color: "#333", padding: "8px 14px", borderRadius: "16px 16px 16px 4px", maxWidth: 280 },
  sender: { fontSize: 11, fontWeight: 600, marginBottom: 4, opacity: 0.7 },
  time: { fontSize: 10, opacity: 0.6, marginTop: 4, textAlign: "right" },
  form: { display: "flex", gap: 8 },
  input: { flex: 1, padding: "12px 16px", border: "1px solid #ddd", borderRadius: 25, fontSize: 14 },
  btn: { background: "#e94560", color: "#fff", border: "none", borderRadius: "50%", width: 48, height: 48, fontSize: 18, cursor: "pointer" },
};