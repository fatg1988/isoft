function search_catalog(catalog_id) {
    localStorage.setItem("blog_search",  JSON.stringify({"catalog_id":catalog_id}));
    window.location.href="/blog/search";
}

function search_text() {
    var search_text = $("input[name='search_text']").val();
    localStorage.setItem("blog_search",  JSON.stringify({"search_text":search_text}));
    window.location.href="/blog/search";
}