#include <linux/module.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/sched.h>
#include <linux/sched/clock.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Angel Sique");
MODULE_DESCRIPTION("Informacion cpu");
MODULE_VERSION("1.0");

static int write_to_proc(struct seq_file *file_proc, void *v)
{
    unsigned long total_cpu_time = jiffies_to_msecs(get_jiffies_64());
    unsigned long total_usage = 0;
    struct task_struct *task;

    // Calculating total CPU time
    for_each_process(task) {
        total_usage += jiffies_to_msecs(task->utime + task->stime);
    }

    // Printing CPU usage information to /proc
    seq_printf(file_proc, "{\n\"cpu_total\":%lu,\n", total_cpu_time);
    seq_printf(file_proc, "\"cpu_porcentaje\":%lu\n", (total_usage * 100) / total_cpu_time);
    seq_printf(file_proc, "}\n");

    return 0;
}

static int abrir_aproc(struct inode *inode, struct file *file)
{
    return single_open(file, write_to_proc, NULL);
}

static struct proc_ops archivo_operaciones = {
    .proc_open = abrir_aproc,
    .proc_read = seq_read
};

static int __init modulo_init(void)
{
    proc_create("cpu_info", 0, NULL, &archivo_operaciones);
    printk(KERN_INFO "Insertando el módulo CPU\n");
    return 0;
}

static void __exit modulo_cleanup(void)
{
    remove_proc_entry("cpu_info", NULL);
    printk(KERN_INFO "Eliminando el módulo CPU\n");
}

module_init(modulo_init);
module_exit(modulo_cleanup);
