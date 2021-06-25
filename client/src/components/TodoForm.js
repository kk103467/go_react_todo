import React, { useState } from 'react'

export default function TodoForm(props) {
    const [taskname, setTaskname] = useState("");

    const handleSubmit = e => {
        e.preventDefault();
        if (!taskname) return;
        props.addTodo(taskname);
        setTaskname("");
    };

    return (
        <form>
            <input
                type="text"
                value={taskname}
                onChange={e => { setTaskname(e.target.value) }}
            />
            <button
                type="submit"
                onClick={handleSubmit}
            >
                Add new task
            </button>
        </form>
    );
}