/** @type {import('svelte/action').Action}  */
export const onDrag = (node, params) => {
    let dragStart = null;

    const attr = params.orientation === "vertical" ? "screenX" : "screenY";

    const mouseDownAction = e => {
        e.preventDefault();
        node.dispatchEvent(new CustomEvent("dragStart", { detail: "hello" }));
        dragStart = e[attr];
    };
    const mouseMoveAction = e => {
        if (dragStart != null) {
            const delta = e[attr] - dragStart;
            node.dispatchEvent(new CustomEvent("drag", { detail: delta }));
        }
    };
    const mouseUpAction = () => {
        dragStart = null;
        node.dispatchEvent(new CustomEvent("dragEnd", { detail: "hello" }));
    };

    node.addEventListener("mousedown", mouseDownAction);
    document.addEventListener("mousemove", mouseMoveAction);
    document.addEventListener("mouseup", mouseUpAction);

    return {
        destroy() {
            node.removeEventListener("mousedown", mouseDownAction);
            document.removeEventListener("mousemove", mouseMoveAction);
            document.removeEventListener("mouseup", mouseUpAction);
        },
    };
};
