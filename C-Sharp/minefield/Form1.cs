using System;
using System.Collections;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Data.OleDb;
using System.Drawing;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Final_Ödevi
{
    public partial class Form1 : Form
    {
        public bool flag = false;
        public Form1()
        {
            InitializeComponent();
        }
        OleDbConnection baglanti;
        OleDbCommand komut;
        OleDbDataAdapter da;
        public void KisiEkle()
        {
            baglanti = new OleDbConnection("Provider=Microsoft.Jet.OLEDB.4.0;Data Source=Veritabanı.mdb;Persist Security Info=False;");
            baglanti.Open();
            da = new OleDbDataAdapter("SELECT *FROM Stok", baglanti);
            DataTable dataTable = new DataTable();
            da.Fill(dataTable);
            dataGridView1.DataSource = dataTable;
            baglanti.Close();
        }

        private void tabPage1_Click(object sender, EventArgs e)
        {

        }

        private void statusStrip1_ItemClicked(object sender, ToolStripItemClickedEventArgs e)
        {

        }

        private void toolStripStatusLabel3_Click(object sender, EventArgs e)
        {

        }

        private void statusStrip1_ItemClicked_1(object sender, ToolStripItemClickedEventArgs e)
        {

        }

        private void toolStripStatusLabel1_Click(object sender, EventArgs e)
        {

        }

        private void timer1_Tick(object sender, EventArgs e)
        {

        }

        private void toolStripStatusLabel2_Click(object sender, EventArgs e)
        {

        }

        private void timer1_Tick_1(object sender, EventArgs e)
        {
            string saat = DateTime.Now.ToLongTimeString();
            string saniye = DateTime.Now.Second.ToString();
           // staProgBar.Value = Convert.ToInt32(saniye);
            toolStripStatusLabel1.Text = "Saat :" + saat;
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            KisiEkle();
            toolStripStatusLabel1.Text = Convert.ToString(DateTime.Now);
        }

        private void button1_Click(object sender, EventArgs e)
        {
            flag = true;
            Form2 form2 = new Form2();
            form2.Show();
            
        }
   

        private void dataGridView1_CellContentClick(object sender, DataGridViewCellEventArgs e)
        {
            dataGridView1.AllowUserToAddRows = false;
            dataGridView1.AllowUserToDeleteRows = false;
            dataGridView1.SelectionMode = DataGridViewSelectionMode.FullRowSelect;
            dataGridView1.ReadOnly = true;
        }

        private void button2_Click(object sender, EventArgs e)
        {
            DialogResult secenek = MessageBox.Show("Kayıt Silinsin mi?", "Dikkat", MessageBoxButtons.YesNo, MessageBoxIcon.Information);

            if (secenek == DialogResult.Yes)
            {
                string sorgu = "DELETE FROM Stok WHERE uadi=@ad ";
                komut = new OleDbCommand(sorgu, baglanti);
                komut.Parameters.AddWithValue("@ad", dataGridView1.CurrentRow.Cells[0].Value);
                baglanti.Open();
                komut.ExecuteNonQuery();
                baglanti.Close();
                KisiEkle();
                MessageBox.Show("Kişi Başarılı Bir Şekilde Silindi");
            }
            else if (secenek == DialogResult.No)
            {
               
            }

        }

        private void button3_Click(object sender, EventArgs e)
        {
            Form2 frm2 = new Form2();
            flag = false;
            frm2.a = dataGridView1.CurrentRow.Cells[0].Value.ToString();
            frm2.b = dataGridView1.CurrentRow.Cells[1].Value.ToString();
            frm2.c = dataGridView1.CurrentRow.Cells[2].Value.ToString();
            frm2.d = dataGridView1.CurrentRow.Cells[3].Value.ToString();
            frm2.Show();
            
        }

        private void button4_Click(object sender, EventArgs e)
        {

            double ab =Convert.ToInt32(dataGridView1.CurrentRow.Cells[1].Value.ToString());
            double ac = Convert.ToInt32(dataGridView1.CurrentRow.Cells[2].Value.ToString());
            
            double satısFiyatı=0;
            satısFiyatı = (ab * (1 + (ac / 100))) * 1.18;

            label1.Text = "Seçilen Ürünün Satış Fiyatı:"+ satısFiyatı+"TL";


        }

        private void toolStripStatusLabel3_Click_1(object sender, EventArgs e)
        {

        }

        private void tabPage2_Click(object sender, EventArgs e)
        {

        }
        ArrayList mayinlar = new ArrayList();
        private void button5_Click(object sender, EventArgs e)
        {
            mayinlar.Clear();
            flowLayoutPanel1.Controls.Clear();
            int mayin1 = 10;
            int mayin2 = 25;
            int mayin3 = 40;
            int tarla = 100;
            int sayi = 0;
            Random random = new Random();
            

            if (radioButton1.Checked) 
            {
                label2.Text = "Mayın Sayısı= 10";
                for (int i = 0; i < mayin1; i++)
                {
                uret:
                    sayi = random.Next(0, tarla);
                    if (mayinlar.Contains(sayi)) 
                    {
                        goto uret;
                    }
                    else 
                    {
                        mayinlar.Add(sayi);
                    }
                }
                for (int i = 0; i < tarla; i++)
                {
                    Button button = new Button();
                    button.Size = new Size(40,40);
                    button.UseVisualStyleBackColor = true;
                    if (mayinlar.Contains(i)) 
                    {
                        button.Tag = -1;
                    }
                    else 
                    {
                        button.Tag = random.Next(1, 10);
                    }
                    button.Click += Button_Click1;
                    flowLayoutPanel1.Controls.Add(button);
                }

            }
            else if (radioButton2.Checked) 
            {
                label2.Text = "Mayın Sayısı= 25";
                for (int i = 0; i < mayin2; i++)
                {
                uret:
                    sayi = random.Next(0, tarla);
                    if (mayinlar.Contains(sayi))
                    {
                        goto uret;
                    }
                    else
                    {
                        mayinlar.Add(sayi);
                    }
                }
                for (int i = 0; i < tarla; i++)
                {
                    Button button = new Button();
                    button.Size = new Size(40, 40);
                    button.BackColor = Color.Blue;
                    button.UseVisualStyleBackColor = true;
                    if (mayinlar.Contains(i))
                    {
                        button.Tag = -1;
                    }
                    else
                    {
                        button.Tag = random.Next(1, 25);
                    }
                    button.Click += Button_Click1;
                    flowLayoutPanel1.Controls.Add(button);
                }
            }
            else if (radioButton3.Checked) 
            {
                label2.Text = "Mayın Sayısı= 40";
                for (int i = 0; i < mayin3; i++)
                {
                uret:
                    sayi = random.Next(0, tarla);
                    if (mayinlar.Contains(sayi))
                    {
                        goto uret;
                    }
                    else
                    {
                        mayinlar.Add(sayi);
                    }
                }
                for (int i = 0; i < tarla; i++)
                {
                    Button button = new Button();
                    button.Size = new Size(40, 40);
                    button.UseVisualStyleBackColor = true;
                    if (mayinlar.Contains(i))
                    {
                        button.Tag = -1;
                    }
                    else
                    {
                        button.Tag = random.Next(1, 40);
                    }
                    button.Click += Button_Click1;
                    flowLayoutPanel1.Controls.Add(button);
                }
            }
            else 
            {
                MessageBox.Show("Seviye seçiniz........!");
            }
        }
        int puan = 0;
        private void Button_Click1(object sender, EventArgs e)
        {
            Button tiklanan = (Button)sender;
            tiklanan.Text=tiklanan.Tag.ToString();
            if(tiklanan.Text=="")
            {
                if (int.Parse(tiklanan.Tag.ToString()) == -1)
                {
                    tiklanan.BackgroundImage = Image.FromFile("mayin.png");
                    tiklanan.BackColor = Color.Red;
                    for (int i = 0; i < flowLayoutPanel1.Controls.Count; i++)
                    {
                        if (int.Parse(flowLayoutPanel1.Controls[i].Tag.ToString()) == -1)
                        {
                            flowLayoutPanel1.Controls[i].BackgroundImage = Image.FromFile("mayin.png");
                            flowLayoutPanel1.Controls[i].Enabled = false;                           
                        }
                        else
                        {
                            flowLayoutPanel1.Controls[i].Text = flowLayoutPanel1.Controls[i].Tag.ToString();
                        }
                    }
                    MessageBox.Show("Yandınız" +
                                "Toplam Puan:" + puan);
                }
                else
                {
                    puan += int.Parse(tiklanan.Tag.ToString());
                    tiklanan.Text = tiklanan.ToString();
                    label3.Text = "Puan: " + puan;
                }

            }
            
        }

    }
}
