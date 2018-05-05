;(function (window, $, undefined) {
    /*
     * tab切换插件
     * 用例：$('*').createTab();
     */
    $.fn.createTab = function (opt) {
        var def = {
            activeEvt: 'mouseover',
            activeCls: 'cur',
            marginLeft:-1000,           // 切换移动的距离
            time:300,                    // 切换移动的速度
            speed:'slow',
        };
        def = $.extend(def, opt);
        this.each(function () {
            var $this = $(this);
            var timer;
            $this.find('ul.title li').mouseover(def.activeEvt,function(){
                var index = $(this).index(),
                    that = $(this);
                timer = setTimeout(function(){
                    that.addClass('cur').siblings().removeClass('cur');
                    $this.find('div.list').animate({marginLeft: def.marginLeft * index},def.speed);
                },def.time);
            }).mouseout(function(){
                clearTimeout( timer );
            })
        });
    }

})(window, jQuery);
$(function(){
    $(".jyTable").createTab({marginLeft:-1000, time : 300, speed : 'slow'})
});
