// automatically generated by the FlatBuffers compiler, do not modify

package MyGame.Sample;

import java.nio.*;
import java.lang.*;
import java.util.*;
import com.google.flatbuffers.*;

@SuppressWarnings("unused")
public final class Monster extends Table {
  public static void ValidateVersion() { Constants.FLATBUFFERS_1_12_0(); }
  public static Monster getRootAsMonster(ByteBuffer _bb) { return getRootAsMonster(_bb, new Monster()); }
  public static Monster getRootAsMonster(ByteBuffer _bb, Monster obj) { _bb.order(ByteOrder.LITTLE_ENDIAN); return (obj.__assign(_bb.getInt(_bb.position()) + _bb.position(), _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __reset(_i, _bb); }
  public Monster __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public MyGame.Sample.Vec3 pos() { return pos(new MyGame.Sample.Vec3()); }
  public MyGame.Sample.Vec3 pos(MyGame.Sample.Vec3 obj) { int o = __offset(4); return o != 0 ? obj.__assign(o + bb_pos, bb) : null; }
  public short mana() { int o = __offset(6); return o != 0 ? bb.getShort(o + bb_pos) : 150; }
  public short hp() { int o = __offset(8); return o != 0 ? bb.getShort(o + bb_pos) : 100; }
  public String name() { int o = __offset(10); return o != 0 ? __string(o + bb_pos) : null; }
  public ByteBuffer nameAsByteBuffer() { return __vector_as_bytebuffer(10, 1); }
  public ByteBuffer nameInByteBuffer(ByteBuffer _bb) { return __vector_in_bytebuffer(_bb, 10, 1); }
  public String names(int j) { int o = __offset(12); return o != 0 ? __string(__vector(o) + j * 4) : null; }
  public int namesLength() { int o = __offset(12); return o != 0 ? __vector_len(o) : 0; }
  public StringVector namesVector() { return namesVector(new StringVector()); }
  public StringVector namesVector(StringVector obj) { int o = __offset(12); return o != 0 ? obj.__assign(__vector(o), 4, bb) : null; }
  public int inventory(int j) { int o = __offset(16); return o != 0 ? bb.get(__vector(o) + j * 1) & 0xFF : 0; }
  public int inventoryLength() { int o = __offset(16); return o != 0 ? __vector_len(o) : 0; }
  public ByteVector inventoryVector() { return inventoryVector(new ByteVector()); }
  public ByteVector inventoryVector(ByteVector obj) { int o = __offset(16); return o != 0 ? obj.__assign(__vector(o), bb) : null; }
  public ByteBuffer inventoryAsByteBuffer() { return __vector_as_bytebuffer(16, 1); }
  public ByteBuffer inventoryInByteBuffer(ByteBuffer _bb) { return __vector_in_bytebuffer(_bb, 16, 1); }
  public byte color() { int o = __offset(18); return o != 0 ? bb.get(o + bb_pos) : 2; }
  public MyGame.Sample.Weapon weapons(int j) { return weapons(new MyGame.Sample.Weapon(), j); }
  public MyGame.Sample.Weapon weapons(MyGame.Sample.Weapon obj, int j) { int o = __offset(20); return o != 0 ? obj.__assign(__indirect(__vector(o) + j * 4), bb) : null; }
  public int weaponsLength() { int o = __offset(20); return o != 0 ? __vector_len(o) : 0; }
  public MyGame.Sample.Weapon.Vector weaponsVector() { return weaponsVector(new MyGame.Sample.Weapon.Vector()); }
  public MyGame.Sample.Weapon.Vector weaponsVector(MyGame.Sample.Weapon.Vector obj) { int o = __offset(20); return o != 0 ? obj.__assign(__vector(o), 4, bb) : null; }
  public byte equippedType() { int o = __offset(22); return o != 0 ? bb.get(o + bb_pos) : 0; }
  public Table equipped(Table obj) { int o = __offset(24); return o != 0 ? __union(obj, o + bb_pos) : null; }
  public MyGame.Sample.Vec3 path(int j) { return path(new MyGame.Sample.Vec3(), j); }
  public MyGame.Sample.Vec3 path(MyGame.Sample.Vec3 obj, int j) { int o = __offset(26); return o != 0 ? obj.__assign(__vector(o) + j * 12, bb) : null; }
  public int pathLength() { int o = __offset(26); return o != 0 ? __vector_len(o) : 0; }
  public MyGame.Sample.Vec3.Vector pathVector() { return pathVector(new MyGame.Sample.Vec3.Vector()); }
  public MyGame.Sample.Vec3.Vector pathVector(MyGame.Sample.Vec3.Vector obj) { int o = __offset(26); return o != 0 ? obj.__assign(__vector(o), 12, bb) : null; }

  public static void startMonster(FlatBufferBuilder builder) { builder.startTable(12); }
  public static void addPos(FlatBufferBuilder builder, int posOffset) { builder.addStruct(0, posOffset, 0); }
  public static void addMana(FlatBufferBuilder builder, short mana) { builder.addShort(1, mana, 150); }
  public static void addHp(FlatBufferBuilder builder, short hp) { builder.addShort(2, hp, 100); }
  public static void addName(FlatBufferBuilder builder, int nameOffset) { builder.addOffset(3, nameOffset, 0); }
  public static void addNames(FlatBufferBuilder builder, int namesOffset) { builder.addOffset(4, namesOffset, 0); }
  public static int createNamesVector(FlatBufferBuilder builder, int[] data) { builder.startVector(4, data.length, 4); for (int i = data.length - 1; i >= 0; i--) builder.addOffset(data[i]); return builder.endVector(); }
  public static void startNamesVector(FlatBufferBuilder builder, int numElems) { builder.startVector(4, numElems, 4); }
  public static void addInventory(FlatBufferBuilder builder, int inventoryOffset) { builder.addOffset(6, inventoryOffset, 0); }
  public static int createInventoryVector(FlatBufferBuilder builder, byte[] data) { return builder.createByteVector(data); }
  public static int createInventoryVector(FlatBufferBuilder builder, ByteBuffer data) { return builder.createByteVector(data); }
  public static void startInventoryVector(FlatBufferBuilder builder, int numElems) { builder.startVector(1, numElems, 1); }
  public static void addColor(FlatBufferBuilder builder, byte color) { builder.addByte(7, color, 2); }
  public static void addWeapons(FlatBufferBuilder builder, int weaponsOffset) { builder.addOffset(8, weaponsOffset, 0); }
  public static int createWeaponsVector(FlatBufferBuilder builder, int[] data) { builder.startVector(4, data.length, 4); for (int i = data.length - 1; i >= 0; i--) builder.addOffset(data[i]); return builder.endVector(); }
  public static void startWeaponsVector(FlatBufferBuilder builder, int numElems) { builder.startVector(4, numElems, 4); }
  public static void addEquippedType(FlatBufferBuilder builder, byte equippedType) { builder.addByte(9, equippedType, 0); }
  public static void addEquipped(FlatBufferBuilder builder, int equippedOffset) { builder.addOffset(10, equippedOffset, 0); }
  public static void addPath(FlatBufferBuilder builder, int pathOffset) { builder.addOffset(11, pathOffset, 0); }
  public static void startPathVector(FlatBufferBuilder builder, int numElems) { builder.startVector(12, numElems, 4); }
  public static int endMonster(FlatBufferBuilder builder) {
    int o = builder.endTable();
    return o;
  }

  public static final class Vector extends BaseVector {
    public Vector __assign(int _vector, int _element_size, ByteBuffer _bb) { __reset(_vector, _element_size, _bb); return this; }

    public Monster get(int j) { return get(new Monster(), j); }
    public Monster get(Monster obj, int j) {  return obj.__assign(__indirect(__element(j), bb), bb); }
  }
}
