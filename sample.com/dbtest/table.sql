/* psqlでログイン後、\i ./table.sqlって感じでデータを作って下さい。 */
/* 戦国大戦のゲームデータからテーブル構成を作ってます。 */
/* http://www29.atwiki.jp/sengoku-taisen/pages/2217.html */
create table character (
	  no integer not null primary key	/* No. */
	, name text not null				/* 名前 */
	, cost decimal not null				/* コスト */
	, type integer not null				/* 兵種 */
	, attack integer not null			/* 武力 */
	, lead integer not null				/* 統率 */
	, scheme text not null				/* 計略名 */
	, morale int not null				/* 士気（計略を使うためのもの） */
);

create table typemaster (
	  type integer not null primary key 
	, name text not null				/* 兵種名 */
);

insert into typemaster values (1, '鉄砲');
insert into typemaster values (2, '騎馬');
insert into typemaster values (3, '弓');
insert into typemaster values (4, '槍');
insert into typemaster values (5, '足軽');

insert into character values(1, '織田信長', 3.0, 1, 9, 10, '天下布武', 6);
insert into character values(2, '柴田勝家', 3.0, 4, 9, 8, '掛かれ柴田', 6);
insert into character values(3, '滝川一益', 2.5, 1, 8, 6, '撹乱貫通射撃', 4);
insert into character values(4, '佐々成政', 2.0, 1, 7, 7, '母衣衆の采配', 5);
