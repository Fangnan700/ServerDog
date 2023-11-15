// 更新数据
function update() {
    fetch("/update")
        .then(resp => {
            return resp.json();
        })
        .then(data => {
            display(data);
        });
}

// 格式化存储容量
function format_size(size) {
    let value = Number(size);
    const units = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB'];
    let index = 0;
    let k = value;
    if (value >= 1024) {
        while (k > 1024) {
            k = k / 1024;
            index++;
        }
    }
    return `${(k).toFixed(2)}${units[index]}`;
}

// 格式化网速
function format_network_speed(speedInBytesPerSecond) {
    let value = Number(speedInBytesPerSecond);
    const units = ['B/s', 'KB/s', 'MB/s', 'GB/s', 'TB/s', 'PB/s', 'EB/s'];
    let index = 0;
    let k = value;
    if (value >= 1024) {
        while (k > 1024) {
            k = k / 1024;
            index++;
        }
    }
    return `${(k).toFixed(2)}${units[index]}`;
}

// 格式化运行时间
function format_running_time(t) {
    t = parseFloat(t);

    let days = Math.floor(t / (24 * 3600));
    let hours = Math.floor((t % (24 * 3600)) / 3600);
    let minutes = Math.floor((t % 3600) / 60);
    let seconds = Math.floor(t % 60);

    return `${days}天${hours}小时${minutes}分钟${seconds}秒`;
}

// 展示数据
function display(data_list) {
    let contents = "";

    // 遍历数据逐个解析
    for(let data of data_list) {
        let info = {
            // 主机信息
            "host_ip": data["SystemInfo"]["host"]["ip"],
            "host_name": data["SystemInfo"]["host"]["host_name"],
            "host_os": data["SystemInfo"]["host"]["os"],
            "host_platform": data["SystemInfo"]["host"]["platform"],
            "host_platform_family": data["SystemInfo"]["host"]["platform_family"],
            "host_platform_version": data["SystemInfo"]["host"]["platform_version"],
            "host_kernel_arch": data["SystemInfo"]["host"]["kernel_arch"],
            "host_kernel_version": data["SystemInfo"]["host"]["kernel_version"],
            "host_location": data["SystemInfo"]["host"]["location"],
            "host_provider": data["SystemInfo"]["host"]["provider"],
            "host_last_login": data["SystemInfo"]["host"]["last_login"],
            "host_running_time": data["SystemInfo"]["host"]["running_time"],

            // CPU信息
            "cpu_model_name": data["SystemInfo"]["cpu"]["model_name"],
            "cpu_cache_size": data["SystemInfo"]["cpu"]["cache_size"],
            "cpu_max_hz": data["SystemInfo"]["cpu"]["max_hz"],
            "cpu_counts": data["SystemInfo"]["cpu"]["counts"],
            "cpu_used_percent": data["SystemInfo"]["cpu"]["used_percent"]*1.0,
            "cpu_process": data["SystemInfo"]["cpu"]["precess"],
            "cpu_threads": data["SystemInfo"]["cpu"]["threads"],

            // 内存信息
            "mem_total": data["SystemInfo"]["mem"]["total"],
            "mem_available": data["SystemInfo"]["mem"]["available"],
            "mem_free": data["SystemInfo"]["mem"]["free"],
            "mem_cached": data["SystemInfo"]["mem"]["cached"],
            "mem_used_percent": data["SystemInfo"]["mem"]["used_percent"]*1.0,

            // 交换分区信息
            "swap_total": data["SystemInfo"]["swap"]["total"],
            "swap_used": data["SystemInfo"]["swap"]["used"],
            "swap_free": data["SystemInfo"]["swap"]["free"],
            "swap_used_percent": data["SystemInfo"]["swap"]["used_percent"]*1.0,

            // 磁盘信息
            "disk_total": data["SystemInfo"]["disk"]["total"],
            "disk_used": data["SystemInfo"]["disk"]["used"],
            "disk_free": data["SystemInfo"]["disk"]["free"],
            "disk_used_percent": data["SystemInfo"]["disk"]["used_percent"]*100.0,
            "disk_read_bytes": data["SystemInfo"]["disk"]["read_bytes"]*1.0,
            "disk_write_bytes": data["SystemInfo"]["disk"]["write_bytes"]*1.0,

            // 网络信息
            "net_upload_speed": data["SystemInfo"]["net"]["upload_speed"],
            "net_download_speed": data["SystemInfo"]["net"]["download_speed"]
        };
        
        
        /*
        * 格式化单位
        * */
        let cpu_used = info.cpu_used_percent;
        let mem_used = info.mem_used_percent;
        let swap_used = info.swap_used_percent;
        let disk_used = info.disk_used_percent;

        info.host_running_time = format_running_time(info.host_running_time);

        info.cpu_model_name = info.cpu_model_name.replaceAll(" ", "");
        info.cpu_cache_size = (parseFloat(info.cpu_cache_size)/1024/1024).toFixed(0) + 'MB';
        info.cpu_used_percent = info.cpu_used_percent.toFixed(2) + '%';
        info.cpu_max_hz = info.cpu_max_hz + 'MHZ';

        info.mem_total = format_size(info.mem_total);
        info.mem_available = format_size(info.mem_available);
        info.mem_free = format_size(info.mem_free);
        info.mem_cached = format_size(info.mem_cached);
        info.mem_used_percent = info.mem_used_percent.toFixed(2) + '%';

        info.swap_total = format_size(info.swap_total);
        info.swap_used = format_size(info.swap_used);
        info.swap_free = format_size(info.swap_free);
        info.swap_used_percent = info.swap_used_percent.toFixed(2) + '%';

        info.disk_total = format_size(info.disk_total);
        info.disk_free = format_size(info.disk_free);
        info.disk_used = format_size(info.disk_used);
        info.disk_read_bytes = format_size(info.disk_read_bytes);
        info.disk_write_bytes = format_size(info.disk_write_bytes);
        info.disk_used_percent = info.disk_used_percent.toFixed(2) + '%';

        info.net_upload_speed = format_network_speed(info.net_upload_speed);
        info.net_download_speed = format_network_speed(info.net_download_speed);

        contents += `
        <div class="info_item">
            <div class="card card-main">
                <div class="card-body">
                    <div class="info_title">主机信息(${info.host_ip})</div>
                    <div class="info_text">主机名称：${info.host_name}</div>
                    <div class="info_text">操作系统：${info.host_os}</div>
                    <div class="info_text">发行版本：${info.host_platform}</div>
                    <div class="info_text">平台家族：${info.host_platform_family}</div>
                    <div class="info_text">平台版本：${info.host_platform_version}</div>
                    <div class="info_text">内核架构：${info.host_kernel_arch}</div>
                    <div class="info_text">内核版本：${info.host_kernel_version}</div>
                    <div class="info_text">最近登录：${info.host_last_login}</div>
                    <div class="info_text">运行时间：${info.host_running_time}</div>
                    <div class="info_text">服务商：${info.host_provider}</div>
                    <div class="info_text">归属地：${info.host_location}</div>
                    <br>
                    <div class="info_title">CPU信息</div>
                    <div class="info_text">CPU型号：${info.cpu_model_name}</div>
                    <div class="info_text">CPU缓存：${info.cpu_cache_size}</div>
                    <div class="info_text">CPU主频：${info.cpu_max_hz}</div>
                    <div class="info_text">CPU核心数：${info.cpu_counts}</div>
                    <div class="info_text">CPU占用比：${info.cpu_used_percent}</div>
                    <div class="info_text">当前进程数：${info.cpu_process}</div>
                    <div class="info_text">当前线程数：${info.cpu_threads}</div>
                </div>
            </div>
    
            <div class="card card-side">
                <div class="card-body">
                    <div class="info_title">磁盘信息</div>
                    <div class="info_text">磁盘容量：${info.disk_total}</div>
                    <div class="info_text">已用容量：${info.disk_used}</div>
                    <div class="info_text">空闲容量：${info.disk_free}</div>
                    <div class="info_text">读取总量：${info.disk_read_bytes}</div>
                    <div class="info_text">写入总量：${info.disk_write_bytes}</div>
                    <div class="info_text">磁盘占用比：${info.disk_used_percent}</div>
                    <br>
                    <div class="info_title">内存信息</div>
                    <div class="info_text">内存容量：${info.mem_total}</div>
                    <div class="info_text">可用内存：${info.mem_available}</div>
                    <div class="info_text">空闲内存：${info.mem_free}</div>
                    <div class="info_text">缓存占用：${info.mem_cached}</div>
                    <div class="info_text">内存占用比：${info.mem_used_percent}</div>
                    <br>
                    <div class="info_title">交换分区信息</div>
                    <div class="info_text">交换分区容量：${info.swap_total}</div>
                    <div class="info_text">已用交换分区：${info.swap_used}</div>
                    <div class="info_text">空闲交换分区：${info.swap_free}</div>
                    <div class="info_text">交换分区占用比：${info.swap_used_percent}</div>
                </div>
            </div>
    
            <div class="card card-side">
                <div class="card-body">
                
                    <div class="svg_frame">
                        <div class="svg_div">
                            <svg width="150" height="150">
                                <circle class="svg_circle" cx="75" cy="75" r="60" stroke="#e0e0e0" stroke-width="10" fill="none" />
                                <circle class="svg_circle" cx="75" cy="75" r="60" stroke="${cpu_used>90 ? '#ffaaa5' : (cpu_used>60 ? '#ffd3b6' : '#a6d0e4')}" stroke-width="10" fill="none" 
                                        stroke-dasharray="502.65"
                                        stroke-dashoffset=${502.65 * (1 - cpu_used/100)}
                                />
                                <text class="svg_text" x="80" y="85" fill="#6b778c" text-anchor="middle">
                                    <tspan x="85" dy="-0.9em">CPU</tspan>
                                    <tspan x="75" dy="1.3em">${info.cpu_used_percent}</tspan>
                                </text>
                            </svg>
                        </div>
                        
                        <div class="svg_div">
                            <svg width="150" height="150">
                                <circle class="svg_circle" cx="75" cy="75" r="60" stroke="#e0e0e0" stroke-width="10" fill="none" />
                                <circle class="svg_circle" cx="75" cy="75" r="60" stroke="${mem_used>90 ? '#ffaaa5' : (mem_used>60 ? '#ffd3b6' : '#a6d0e4')}" stroke-width="10" fill="none" 
                                        stroke-dasharray="502.65"
                                        stroke-dashoffset=${502.65 * (1 - mem_used/100)}
                                />
                                <text class="svg_text" x="80" y="85" fill="#6b778c" text-anchor="middle">
                                    <tspan x="85" dy="-0.9em">RAM</tspan>
                                    <tspan x="75" dy="1.3em">${info.mem_used_percent}</tspan>
                                </text>
                            </svg>
                        </div>
                    </div>  
                    
                    <div class="svg_frame">
                        <div class="svg_div">
                            <svg width="150" height="150">
                                <circle class="svg_circle" cx="75" cy="75" r="60" stroke="#e0e0e0" stroke-width="10" fill="none" />
                                <circle class="svg_circle" cx="75" cy="75" r="60" stroke="${swap_used>90 ? '#ffaaa5' : (swap_used>60 ? '#ffd3b6' : '#a6d0e4')}" stroke-width="10" fill="none" 
                                        stroke-dasharray="502.65"
                                        stroke-dashoffset=${502.65 * (1 - swap_used/100)}
                                />
                                <text class="svg_text" x="80" y="85" fill="#6b778c" text-anchor="middle">
                                    <tspan x="85" dy="-0.9em">Swap</tspan>
                                    <tspan x="75" dy="1.3em">${info.swap_used_percent}</tspan>
                                </text>
                            </svg>
                        </div>
                        
                        <div class="svg_div">
                            <svg width="150" height="150">
                                <circle class="svg_circle" cx="75" cy="75" r="60" stroke="#e0e0e0" stroke-width="10" fill="none" />
                                <circle class="svg_circle" cx="75" cy="75" r="60" stroke="${disk_used>90 ? '#ffaaa5' : (disk_used>60 ? '#ffd3b6' : '#a6d0e4')}" stroke-width="10" fill="none" 
                                        stroke-dasharray="502.65"
                                        stroke-dashoffset=${502.65 * (1 - disk_used/100)}
                                />
                                <text class="svg_text" x="80" y="85" fill="#6b778c" text-anchor="middle">
                                    <tspan x="85" dy="-0.9em">Disk</tspan>
                                    <tspan x="75" dy="1.3em">${info.disk_used_percent}</tspan>
                                </text>
                            </svg>
                        </div>
                    </div>
                    
                    <div class="svg_frame">
                        <div class="info_text net_info">上行：${info.net_upload_speed}</div>
                        <div class="info_text net_info">下行：${info.net_download_speed}</div>
                    </div>
                     
                </div>
            </div>
        </div>
        `
    }

    document.getElementById("content").innerHTML = contents;
}

function add_server() {
    let server_addr = document.getElementById("server_addr_input").value;

    document.getElementById("server_addr_input").value = "";
    document.getElementById("server_addr_btn").blur();
    console.log(server_addr);

    let options = {
        method: "POST",
        body: JSON.stringify({
            "addr": server_addr
        })
    };

    fetch("/add",options)
        .then(resp => {
            let toastElement = document.getElementById("liveToast");
            let toastMessage = document.getElementById("messageText");

            if(resp.status === 200) {
                toastMessage.innerHTML = `<svg t="1700024911398" class="icon" viewBox="0 0 1479 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5784" width="32" height="32"><path d="M1401.287111 0L1479.111111 77.824 544.938667 1012.053333 0 467.114667l136.248889-136.248889 447.601778 291.896889L1401.287111 0z" fill="#a6d0e4" p-id="5785"></path></svg>&nbsp&nbsp<span class="message_text">添加主机成功！</span>`
            } else {
                toastMessage.innerHTML = `<svg t="1700025309866" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="7452" width="32" height="32"><path d="M512 0a512 512 0 0 0-512 512 512 512 0 0 0 512 512 512 512 0 0 0 512-512 512 512 0 0 0-512-512z" fill="#ffaaa5" p-id="7453" data-spm-anchor-id="a313x.search_index.0.i20.6e8f3a818LU6ec" class="selected"></path><path d="M513.755429 565.540571L359.277714 720.018286a39.058286 39.058286 0 0 1-55.296-0.073143 39.277714 39.277714 0 0 1 0.073143-55.442286l154.331429-154.331428-155.062857-155.136a36.571429 36.571429 0 0 1 51.712-51.785143l365.714285 365.714285a36.571429 36.571429 0 1 1-51.785143 51.785143L513.755429 565.540571z m157.549714-262.582857a35.254857 35.254857 0 1 1 49.737143 49.737143l-106.057143 108.982857a35.254857 35.254857 0 1 1-49.883429-49.810285l106.203429-108.982858z" fill="#FFFFFF" p-id="7454"></path></svg>&nbsp&nbsp<span class="message_text">添加主机失败！</span>`
            }

            let toast = new bootstrap.Toast(toastElement);
            toast.show();
        })
}

setInterval(update, 1200);