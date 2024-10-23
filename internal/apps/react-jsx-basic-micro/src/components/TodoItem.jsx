import PropTypes from "prop-types";
import { Trash2 } from "lucide-react";

TodoItem.propTypes = {
  item: PropTypes.object.isRequired,
  onCompletedChange: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default function TodoItem({ item, onCompletedChange, onDelete }) {
  return (
    <div className="flex items-center gap-1">
      <label className="flex items-center gap-2 border rounded-md p-2 border-gray-400 bg-white hover:bg-slate-50 grow">
        <input
          type="checkbox"
          checked={item.completed}
          onChange={(e) => onCompletedChange(item.id, e.target.checked)}
          className="scale-125"
        />
        <span className={item.completed ? "line-through text-gray-400" : ""}>
          {item.title}
        </span>
      </label>
      <button onClick={() => onDelete(item.id)} className="p-2">
        <Trash2 size={20} className="text-gray-500" />
      </button>
    </div>
  );
}
