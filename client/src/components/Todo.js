export default function Todo(props) {
    const handleClickComplete = () => {
      props.changeComplete(props.todo.id)
    };
  
    const handleClickDelete = () => {
      props.deleteTodo(props.todo.id);
    }
  
    const handleTextChange = (e) => {
      props.changeText(props.todo.id, e.target.value)
    }
  
    return (
      <div className="todo"
        style={{ textDecoration: props.todo.isCompleted ? "line-through" : "" }}
      >
        <input
          type="text"
          value={props.todo.text}
          onChange={handleTextChange}
        ></input>
        <div>
          <button onClick={handleClickComplete}>Change Complete Condition</button>
          <button onClick={handleClickDelete}>x</button>
        </div>
      </div>
    );
  }