<script>
    // @ts-nocheck
    import Sidebar from "/src/components/Sidebar/Sidebar.svelte";
    import Footer from "/src/components/Footer.svelte";
    import { onDrag } from "/src/components/Sidebar/SidebarResizer";
    import Header from "/src/components/Header.svelte";
    import Messages from "/src/components/Chatbox.svelte";
    import Chatbox from "/src/components/Chatbox.svelte";
    
    let width = 205;
    let resizeWidth = width;
    const minWidth = 205;
    const maxWidth = () => window.innerWidth / 2;

    $: console.log(resizeWidth);

    function handleDrag(delta) {
        const newWidth = width + delta;
        resizeWidth = Math.max(minWidth, Math.min(maxWidth(), newWidth));
    }

    let footerHeight = 96; // Default footer height in pixels (24px * 4)
    let resizeFooterHeight = footerHeight;
    const minFooterHeight = 96; // Minimum footer height in pixels
    const maxFooterHeight = () => window.innerHeight / 2; // Maximum footer height

    function handleFooterDrag(delta) {
        const newHeight = footerHeight - delta;
        resizeFooterHeight = Math.max(minFooterHeight, Math.min(maxFooterHeight(), newHeight));
    }
</script>

<main class="h-full select-none">
    <div class="wails-drag w-full h-6 m-0"></div>
    <div class="flex w-full h-[calc(100vh-1.5rem)]">
        <div
            class="min-w-[215px] max-w-[50%] mt-1"
            style="width: {resizeWidth}px;"
        >
            <Sidebar />
        </div>
        <div
            role="separator"
            aria-orientation="vertical"
            class="w-2 min-w-2 cursor-col-resize"
            use:onDrag={{ orientation: "vertical" }}
            on:drag={({ detail: delta }) => handleDrag(delta)}
            on:dragEnd={() => {
                width = resizeWidth;
            }}
        ></div>
        <div class="flex flex-col flex-grow mr-2 mb-2">
            <div class="flex-none h-16 rounded-md mt-1">
                <Header recipient={"Some Person"} />
            </div>
                <Chatbox />
            <div
                role="separator"
                aria-orientation="horizontal"
                class="h-2 min-h-2 cursor-row-resize"
                use:onDrag={{ orientation: "horizontal" }}
                on:drag={({ detail: delta }) => handleFooterDrag(delta)}
                on:dragEnd={() => {
                    footerHeight = resizeFooterHeight;
                }}
            ></div>
            <div
                class="flex-none rounded-md"
                style="height: {resizeFooterHeight}px;"
            >
                <Footer recipient={"Some Person"} />
            </div>
        </div>
    </div>
</main>
