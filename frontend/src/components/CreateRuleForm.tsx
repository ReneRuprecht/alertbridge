import { useState } from "react";
import type { CreateRuleRequest } from "../types/Rule";
import { createRule } from "../api/CreateRule";
import './CreateRuleForm.css'

interface RuleFormProps {
    onSuccess: Function;
}

export default function CreateRuleForm({ onSuccess }: RuleFormProps) {
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const [priority, setPriority] = useState(1);
    const [enabled, setEnabled] = useState(true);
    const [loading, setLoading] = useState(false);

    const handleSubmit = async (e: any) => {
        e.preventDefault();

        setLoading(true);

        try {
            const createRuleRequest: CreateRuleRequest = {
                name: name,
                description: description,
                priority: Number(priority),
                enabled: enabled
            }
            const res = await createRule(createRuleRequest)

            if (res.ok) {
                onSuccess()
            }

            setName("");
            setDescription("");
            setPriority(1);
            setEnabled(true);
        } catch (err) {
            console.error("Error creating rule:", err);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="create-rule-wrapper">
            <h2>Regel Anlegen</h2>

            <form onSubmit={handleSubmit}>

                <div>
                    <label>Name</label>
                    <input
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        required
                    />
                </div>

                <div>
                    <label>Description</label>
                    <input
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                    />
                </div>

                <div>
                    <label>Priority</label>
                    <input
                        type="number"
                        value={priority}
                        onChange={(e) => setPriority(Number(e.target.value))}
                    />
                </div>

                <div>
                    <label>
                        <input
                            type="checkbox"
                            checked={enabled}
                            onChange={(e) => setEnabled(e.target.checked)}
                        />
                        Enabled
                    </label>
                </div>

                <button type="submit" disabled={loading}>
                    {loading ? "Creating..." : "Create Rule"}
                </button>
            </form>
        </div>
    );
}
